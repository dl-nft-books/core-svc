package handlers

import (
	"gitlab.com/tokend/nft-books/generator-svc/resources"
	"net/http"

	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/responses"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/helpers"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/requests"
)

func ValidatePromocodeById(w http.ResponseWriter, r *http.Request) {
	logger := helpers.Log(r)

	request, err := requests.NewValidatePromocodeRequest(r)
	if err != nil {
		logger.WithError(err).Error("failed to fetch validate promocode request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	promocode, err := helpers.DB(r).Promocodes().FilterByPromocode(request.Promocode).Get()
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
