package handlers

import (
	"fmt"
	bookModel "github.com/dl-nft-books/book-svc/connector/models"
	bookResourse "github.com/dl-nft-books/book-svc/resources"
	"github.com/dl-nft-books/core-svc/internal/data"
	"github.com/dl-nft-books/core-svc/internal/service/api/helpers"
	"github.com/dl-nft-books/core-svc/internal/service/api/requests"
	"github.com/dl-nft-books/core-svc/internal/service/api/responses"
	"github.com/dl-nft-books/core-svc/internal/signature"
	"github.com/dl-nft-books/core-svc/resources"
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3"
	"math"
	"math/big"
	"net/http"
	"time"
)

// discount in contract is int number where 1% = 10^25
// discount in database is float number where 1% = 0.01
var (
	formattedDiscountMultiplier, _ = big.NewInt(0).SetString(fmt.Sprintf("10000000000000000000000000"), 10)
	ok                             bool
)

func SignMint(w http.ResponseWriter, r *http.Request) {
	logger := helpers.Log(r)

	request, err := requests.NewSignMintRequest(r)
	if err != nil {
		logger.WithError(err).Error("failed to fetch new sign mint request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	// Getting task's mintInfo
	task, err := helpers.DB(r).Tasks().GetById(request.TaskID)
	if err != nil {
		logger.WithError(err).Error("failed to get task")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if task == nil {
		ape.RenderErr(w, problems.NotFound())
		return
	}

	// Getting book's mintInfo
	bookResponse, err := helpers.Booker(r).ListBooks(bookModel.ListBooksParams{
		Id:      []int64{task.BookId},
		ChainId: []int64{task.ChainId},
	})
	if err != nil {
		logger.WithError(err).Error("failed to get a book")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if len(bookResponse.Data) == 0 {
		logger.Warnf("Book with specified id %d in network with chain id %d was not found", task.BookId, task.ChainId)
		ape.RenderErr(w, problems.NotFound())
		return
	}
	book := bookResponse.Data[0]
	// Forming signature mintInfo
	mintConfig := helpers.Minter(r)

	domainData := signature.EIP712DomainData{
		VerifyingAddress: book.Attributes.Networks[0].Attributes.ContractAddress,
		//ContractName:     book.Attributes.ContractName,
		//ContractVersion:  book.Attributes.ContractVersion,
		ChainID: book.Attributes.Networks[0].Attributes.ChainId,
	}
	mintInfo := signature.MintInfo{
		TokenAddress: request.TokenAddress,
		TokenURI:     task.MetadataIpfsHash,
		EndTimestamp: time.Now().Add(mintConfig.Expiration).Unix(),
	}

	mintInfo.PricePerOneToken, err = getPricePerOneToken(w, r, request, book, mintConfig.Precision)
	if err != nil {
		logger.WithError(err).Error("failed to get price")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	// Getting promocode info
	promocode, err := helpers.DB(r).Promocodes().FilterById(request.PromocodeID).Get()
	if err != nil {
		logger.WithError(err).Error("failed to get promocode")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	//isVoucherTokenApplied := book.Data.Attributes.VoucherToken == request.TokenAddress
	//
	//mintInfo.Discount, ok = getPromocodeDiscount(w, r, isVoucherTokenApplied, promocode)
	//if !ok {
	//	return
	//}

	// Signing the mint transaction
	mintSignature, err := signature.SignMintInfo(&mintInfo, &domainData, &mintConfig)
	if err != nil {
		logger.WithError(err).Error("failed to generate eip712 mint signature")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	// Using promocode after signature is formed
	if promocode != nil {
		if err = helpers.DB(r).Promocodes().New().UpdateUsages(promocode.Usages + 1).FilterUpdateById(promocode.Id).Update(); err != nil {
			logger.WithError(err).WithFields(logan.F{"promocode": promocode.Promocode}).Error("failed to update promocode")
			ape.RenderErr(w, problems.InternalError())
			return
		}

		logger.Info("promocode applied, discount: ", mintInfo.Discount.String())
	}

	ape.Render(w, responses.NewSignMintResponse(
		mintInfo.PricePerOneToken.String(),
		mintInfo.Discount.String(),
		mintSignature,
		mintInfo.EndTimestamp,
	))
}

func getPromocodeDiscount(w http.ResponseWriter, r *http.Request, isVoucherTokenApplied bool, promocode *data.Promocode) (*big.Int, bool) {
	logger := helpers.Log(r)
	// Promocodes and vouchers can't be used together
	if !isVoucherTokenApplied && promocode != nil {
		//Validating promocode
		promocodeResponse, err := responses.NewValidatePromocodeResponse(*promocode)

		if err != nil {
			logger.WithError(err).Error("failed to get promocode response")
			ape.RenderErr(w, problems.InternalError())
			return nil, false
		}

		//Checking promocode state
		if promocodeResponse.Data.Attributes.State != resources.PromocodeActive {
			logger.WithError(err).Error(Inactive())
			ape.RenderErr(w, problems.BadRequest(Inactive())...)
			return nil, false
		}

		//Calculating discount
		discount := big.NewInt(int64(promocode.Discount * math.Pow10(helpers.Promocoder(r).Decimal)))

		discount.Mul(discount, formattedDiscountMultiplier)

		return discount, true
	}

	//No discount applied
	return big.NewInt(0), true
}
func getPricePerOneToken(w http.ResponseWriter, r *http.Request, request *requests.SignMintRequest, book bookResourse.Book, precision int) (*big.Int, error) {
	logger := helpers.Log(r)
	//if book.Data.Attributes.VoucherToken == request.TokenAddress {
	//	return big.NewInt(0), nil
	//}

	// Normal scenario without voucher
	// Getting price per token in dollars
	priceResponse, err := helpers.Pricer(r).GetPrice(request.Platform, request.TokenAddress, book.Attributes.Networks[0].Attributes.ChainId)
	if err != nil {
		logger.WithError(err).Error("failed to get price response")
		ape.RenderErr(w, problems.InternalError())
		return nil, errors.Wrap(err, "failed to get price response")
	}
	// Converting price
	return helpers.ConvertPrice(priceResponse.Data.Attributes.Price, precision)
}
