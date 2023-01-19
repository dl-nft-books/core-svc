package handlers

import (
	"fmt"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/tokend/nft-books/generator-svc/internal/data"
	"gitlab.com/tokend/nft-books/generator-svc/resources"
	"log"
	"math"
	"math/big"
	"net/http"
	"time"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/helpers"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/requests"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/responses"
	"gitlab.com/tokend/nft-books/generator-svc/internal/signature"
)

// discount in contract is int number where 1% = 10^25
// discount in database is float number where 1% = 0.01
const discountMultiplier = "10000000000000000000000000" // 10^25

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
	book, err := helpers.Booker(r).GetBookById(task.BookId)
	if err != nil {
		logger.WithError(err).Error("failed to get a book")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if book == nil {
		logger.Warnf("Book with specified id %d was not found", task.BookId)
		ape.RenderErr(w, problems.NotFound())
		return
	}

	// Forming signature mintInfo
	mintConfig := helpers.Minter(r)

	domainData := signature.EIP712DomainData{
		VerifyingAddress: book.Data.Attributes.ContractAddress,
		ContractName:     book.Data.Attributes.ContractName,
		ContractVersion:  book.Data.Attributes.ContractVersion,
		ChainID:          book.Data.Attributes.ChainId,
	}
	log.Println(book.Data.Attributes.ChainId)
	mintInfo := signature.MintInfo{
		TokenAddress: request.TokenAddress,
		TokenURI:     task.MetadataIpfsHash,
	}

	isVoucherTokenApplied := book.Data.Attributes.VoucherToken == request.TokenAddress
	// If using voucher token --> setting price to 0
	if isVoucherTokenApplied {
		mintInfo.PricePerOneToken = big.NewInt(0)
	}

	if !isVoucherTokenApplied {
		// Normal scenario without voucher
		// Getting price per token in dollars
		priceResponse, err := helpers.Pricer(r).GetPrice(request.Platform, request.TokenAddress)
		if err != nil {
			logger.WithError(err).Error("failed to get price")
			ape.RenderErr(w, problems.InternalError())
			return
		}
		// Converting price
		mintInfo.PricePerOneToken, err = helpers.ConvertPrice(priceResponse.Data.Attributes.Price, mintConfig.Precision)
		if err != nil {
			logger.WithError(err).Error("failed to convert price")
			ape.RenderErr(w, problems.InternalError())
			return
		}
	}

	mintInfo.EndTimestamp = time.Now().Add(mintConfig.Expiration).Unix()

	// Getting promocode info
	promocode, err := helpers.DB(r).Promocodes().FilterById(request.PromocodeID).Get()
	if err != nil {
		logger.WithError(err).Error("failed to get promocode")
		ape.RenderErr(w, problems.InternalError())
	}

	// Promocodes and vouchers can't be used together
	if isVoucherTokenApplied {
		mintInfo.Discount = big.NewInt(0)
	}

	if !isVoucherTokenApplied {
		discount, err := getPromocodeDiscount(w, r, promocode)

		if err != nil {
			logger.WithError(err).Error("failed to get discount")
			return
		}

		mintInfo.Discount = discount
	}

	// Signing the mint transaction
	mintSignature, err := signature.SignMintInfo(&mintInfo, &domainData, &mintConfig)
	if err != nil {
		logger.WithError(err).Error("failed to generate eip712 mint signature")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	// Using promocode after signature is formed
	if promocode != nil && !isVoucherTokenApplied {
		if err = helpers.DB(r).Promocodes().New().UpdateUsages(promocode.Usages + 1).FilterById(promocode.Id).Update(); err != nil {
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

var formattedDiscountMultiplier, _ = big.NewInt(0).SetString(fmt.Sprintf(discountMultiplier), 10)

func getPromocodeDiscount(w http.ResponseWriter, r *http.Request, promocode *data.Promocode) (*big.Int, error) {
	if promocode != nil {
		//Validating promocode
		promocodeResponse, err := responses.NewValidatePromocodeResponse(*promocode)

		if err != nil {
			ape.RenderErr(w, problems.InternalError())
			return nil, errors.Wrap(err, "failed to get promocode response")
		}

		//Checking promocode state
		if promocodeResponse.Data.Attributes.State != resources.PromocodeActive {
			expiredError := errors.New(fmt.Sprintf("promocode with state %v is invalid",
				promocodeResponse.Data.Attributes.State))
			ape.RenderErr(w, problems.BadRequest(expiredError)...)
			return nil, errors.Wrap(err, expiredError.Error())
		}

		//Calculating discount
		discount := big.NewInt(int64(promocode.Discount * math.Pow10(helpers.Promocoder(r).Decimal)))

		discount.Mul(discount, formattedDiscountMultiplier)

		return discount, nil
	}

	//No discount applied
	return big.NewInt(0), nil
}
