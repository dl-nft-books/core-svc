package handlers

import (
	"net/http"
	"time"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/helpers"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/requests"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/responses"
)

func GetPrice(w http.ResponseWriter, r *http.Request) {
	logger := helpers.Log(r)

	req, err := requests.NewGetPriceRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	// getting task info
	task, err := helpers.GeneratorDB(r).Tasks().GetById(req.TaskID)
	if err != nil {
		logger.WithError(err).Debug("failed to get task")
		ape.Render(w, problems.InternalError())
		return
	}
	if task == nil {
		ape.Render(w, problems.NotFound())
		return
	}

	// getting book info
	book, err := helpers.BooksQ(r).FilterActual().FilterByID(task.BookId).Get()
	if err != nil {
		ape.Render(w, problems.InternalError())
		return
	}
	if book == nil {
		ape.Render(w, problems.NotFound())
		return
	}

	// getting price in $
	priceRes, err := helpers.Pricer(r).GetPrice(req.Platform, req.TokenAddress)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get price")
		ape.Render(w, problems.InternalError())
		return
	}

	// forming signature info
	mintConfig := helpers.Minter(r)

	info := helpers.SignInfo{
		ContractAddress: book.ContractAddress,
		ContractName:    book.ContractName,
		ContractVersion: book.ContractVersion,
		TokenAddress:    req.TokenAddress,
		ChainID:         mintConfig.ChainID,
		TokenURI:        task.IpfsHash,
	}

	info.Price, err = helpers.ConvertPrice(priceRes.Data.Attributes.Price, mintConfig.Precision)
	if err != nil {
		ape.Render(w, problems.InternalError())
		return
	}

	info.EndTimestamp = time.Now().Add(mintConfig.Expiration).Unix()

	// signing
	signature, err := helpers.Sign(&info, &mintConfig)
	if err != nil {
		ape.Render(w, problems.InternalError())
		return
	}

	ape.Render(w, responses.NewGetPriceResponse(info.Price.String(), signature, info.EndTimestamp))
}
