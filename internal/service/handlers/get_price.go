package handlers

import (
	"net/http"

	"gitlab.com/tokend/nft-books/generator-svc/internal/service/responses"

	"gitlab.com/tokend/nft-books/generator-svc/internal/service/helpers"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/requests"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetPrice(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewGetPriceRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	book, err := helpers.BooksQ(r).FilterByID(req.ID).Get()
	if err != nil {
		ape.Render(w, problems.InternalError())
		return
	}
	if book == nil {
		ape.Render(w, problems.NotFound())
		return
	}

	signature, err := helpers.Sign(r, book.Price, book.ContractAddress)
	if err != nil {
		ape.Render(w, problems.InternalError())
		return
	}

	ape.Render(w, responses.NewGetPriceResponse(book.Price, signature))
}
