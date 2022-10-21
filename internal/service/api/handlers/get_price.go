package handlers

import (
	"fmt"
	helpers2 "gitlab.com/tokend/nft-books/generator-svc/internal/service/api/helpers"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/requests"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/responses"
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetPrice(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewGetPriceRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	book, err := helpers2.DB(r).Books().FilterActual().FilterByID(req.ID).Get()
	if err != nil {
		ape.Render(w, problems.InternalError())
		return
	}
	if book == nil {
		ape.Render(w, problems.NotFound())
		return
	}

	mintConfig := helpers2.Minter(r)

	info := helpers2.SignInfo{
		ContractAddress: book.ContractAddress,
		ContractName:    book.ContractName,
		ContractVersion: book.ContractVersion,
		TokenAddress:    req.TokenAddress,
		ChainID:         mintConfig.ChainID,
	}

	rawPrice, err := helpers2.GetPrice(r, info.TokenAddress, req.Platform)
	if err != nil {
		ape.Render(w, problems.InternalError())
		return
	}

	info.Price, err = helpers2.ConvertPrice(rawPrice, mintConfig.Precision)
	if err != nil {
		ape.Render(w, problems.InternalError())
		return
	}

	signature, err := helpers2.Sign(&info, &mintConfig)
	if err != nil {
		fmt.Println(err)
		ape.Render(w, problems.InternalError())
		return
	}

	ape.Render(w, responses.NewGetPriceResponse(info.Price.String(), signature))
}
