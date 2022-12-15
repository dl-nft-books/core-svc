package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/helpers"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/requests"
)

func DeletePromocodeById(w http.ResponseWriter, r *http.Request) {
	logger := helpers.Log(r)

	request, err := requests.NewPromocodeByIdRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	promocode, err := helpers.DB(r).Promocodes().FilterById(request.Id).Get()
	if err != nil {
		logger.WithError(err).Error("failed to get promocode")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if promocode == nil {
		ape.RenderErr(w, problems.NotFound())
		return
	}

	err = helpers.DB(r).Promocodes().FilterById(request.Id).Delete()
	w.WriteHeader(http.StatusNoContent)
}
