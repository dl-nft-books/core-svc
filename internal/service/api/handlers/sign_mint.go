package handlers

import (
	"net/http"
	"time"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/helpers"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/requests"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/responses"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/signature"
)

func SignMint(w http.ResponseWriter, r *http.Request) {
	logger := helpers.Log(r)

	req, err := requests.NewSignMintRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	// getting task mintInfo
	task, err := helpers.GeneratorDB(r).Tasks().GetById(req.TaskID)
	if err != nil {
		logger.WithError(err).Debug("failed to get task")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if task == nil {
		ape.RenderErr(w, problems.NotFound())
		return
	}

	// getting book mintInfo
	book, err := helpers.BooksQ(r).FilterActual().FilterByID(task.BookId).Get()
	if err != nil {
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if book == nil {
		ape.RenderErr(w, problems.NotFound())
		return
	}

	// getting price in $
	priceRes, err := helpers.Pricer(r).GetPrice(req.Platform, req.TokenAddress)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get price")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	// forming signature mintInfo
	mintConfig := helpers.Minter(r)

	domainData := signature.EIP712DomainData{
		VerifyingAddress: book.ContractAddress,
		ContractName:     book.ContractName,
		ContractVersion:  book.ContractVersion,
		ChainID:          mintConfig.ChainID,
	}

	mintInfo := signature.MintInfo{
		TokenAddress: req.TokenAddress,
		TokenURI:     task.MetadataIpfsHash,
	}

	mintInfo.PricePerOneToken, err = helpers.ConvertPrice(priceRes.Data.Attributes.Price, mintConfig.Precision)
	if err != nil {
		ape.RenderErr(w, problems.InternalError())
		return
	}

	mintInfo.EndTimestamp = time.Now().Add(mintConfig.Expiration).Unix()

	// signing
	signature, err := signature.SignMintInfo(&mintInfo, &domainData, &mintConfig)
	if err != nil {
		logger.WithError(err).Debug("failed to generate eip712 mint signature")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, responses.NewSignMintResponse(
		mintInfo.PricePerOneToken.String(),
		signature,
		mintInfo.EndTimestamp,
	))
}
