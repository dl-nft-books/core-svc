package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/helpers"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/requests"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/responses"
)

func GetPrice(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewGetPriceRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	book, err := helpers.BooksQ(r).FilterActual().FilterByID(req.BookID).Get()
	if err != nil {
		ape.Render(w, problems.InternalError())
		return
	}
	if book == nil {
		ape.Render(w, problems.NotFound())
		return
	}

	mintConfig := helpers.Minter(r)

	info := helpers.SignInfo{
		ContractAddress: book.ContractAddress,
		ContractName:    book.ContractName,
		ContractVersion: book.ContractVersion,
		TokenAddress:    req.TokenAddress,
		ChainID:         mintConfig.ChainID,
		TokenURI:        req.TokenURI,
	}

	priceRes, err := helpers.Pricer(r).GetPrice(req.Platform, req.TokenAddress)
	if err != nil {
		helpers.Log(r).WithError(err).Error("error")
		ape.Render(w, problems.InternalError())
		return
	}

	info.Price, err = helpers.ConvertPrice(priceRes.Data.Attributes.Price, mintConfig.Precision)
	if err != nil {
		ape.Render(w, problems.InternalError())
		return
	}

	signature, err := helpers.Sign(&info, &mintConfig)
	if err != nil {
		ape.Render(w, problems.InternalError())
		return
	}

	ape.Render(w, responses.NewGetPriceResponse(info.Price.String(), signature))
}
