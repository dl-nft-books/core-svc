package handlers

import (
	"github.com/dl-nft-books/core-svc/resources"
	"net/http"

	"github.com/dl-nft-books/core-svc/internal/service/api/responses"

	"github.com/dl-nft-books/core-svc/internal/service/api/helpers"
	"github.com/dl-nft-books/core-svc/internal/service/api/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func ValidatePromocodeById(w http.ResponseWriter, r *http.Request) {
	logger := helpers.Log(r)

	request, err := requests.NewValidatePromocodeRequest(r)
	if err != nil {
		logger.WithError(err).Error("failed to fetch validate promocode request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	promocode, err := helpers.DB(r).Promocodes().FilterByPromocode(request.Promocode).FilterByBookId(request.Bookid).Get()
	if err != nil {
		logger.WithError(err).Error("failed to get promocode")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	var promocodeResponse *resources.ValidatePromocodeResponse

	if promocode == nil {
		ape.RenderErr(w, problems.NotFound())
		return
	}
	promocodeResponse, err = responses.NewValidatePromocodeResponse(*promocode)

	if err != nil {
		logger.WithError(err).Error("failed to get promocode response")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, promocodeResponse)
}
