package handlers

import (
	"fmt"
	bookModel "github.com/dl-nft-books/book-svc/connector/models"
	"math/big"
	"net/http"
	"time"

	"github.com/dl-nft-books/core-svc/internal/service/api/helpers"
	"github.com/dl-nft-books/core-svc/internal/service/api/requests"
	"github.com/dl-nft-books/core-svc/internal/service/api/responses"
	"github.com/dl-nft-books/core-svc/internal/signature"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func SignMintByNft(w http.ResponseWriter, r *http.Request) {
	logger := helpers.Log(r)

	request, err := requests.NewSignMintByNftRequest(r)
	if err != nil {
		logger.WithError(err).Error("failed to fetch new sign mint by nft request")
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
		ChainID:          book.Attributes.Networks[0].Attributes.ChainId,
	}
	mintInfo := signature.MintInfo{
		TokenAddress: request.NftAddress,
		Discount:     big.NewInt(0),
		TokenURI:     task.MetadataIpfsHash,
	}

	// Getting price per token in dollars
	priceResponse, err := helpers.Pricer(r).GetNftPrice(request.Platform, request.NftAddress)
	if err != nil {
		logger.WithError(err).Error("failed to get nft floor price")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	// Converting price
	mintInfo.PricePerOneToken, err = helpers.ConvertPrice(fmt.Sprintf("%f", priceResponse.Data.Attributes.Usd), mintConfig.Precision)
	if err != nil {
		logger.WithError(err).Error("failed to convert price")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	mintInfo.EndTimestamp = time.Now().Add(mintConfig.Expiration).Unix()

	// Signing the mint transaction
	mintSignature, err := signature.SignMintInfo(&mintInfo, &domainData, &mintConfig)
	if err != nil {
		logger.WithError(err).Error("failed to generate eip712 mint signature")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, responses.NewSignMintResponse(
		mintInfo.PricePerOneToken.String(),
		mintInfo.Discount.String(),
		mintSignature,
		mintInfo.EndTimestamp,
	))
}
