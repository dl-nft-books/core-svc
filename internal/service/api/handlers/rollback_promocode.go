package handlers

import (
	"gitlab.com/tokend/nft-books/generator-svc/resources"
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/helpers"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/requests"
)

func RollbackPromocode(w http.ResponseWriter, r *http.Request) {
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
	if promocode.State == resources.PromocodeExpired {
		logger.WithError(err).Info("promocode has been expired")
		errorInactive := problems.Forbidden()
		errorInactive.Detail = "promocode has been expired"
		ape.RenderErr(w, errorInactive)
		return
	}

	promocodesQ := helpers.DB(r).Promocodes().New().UpdateLeftUsages(promocode.LeftUsages + 1)

	if promocode.LeftUsages == 0 {
		promocodesQ = promocodesQ.UpdateState(resources.PromocodeActive)
	}

	err = promocodesQ.Update(promocode.Id)

	if err != nil {
		logger.WithError(err).Error("failed to get promocode")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
