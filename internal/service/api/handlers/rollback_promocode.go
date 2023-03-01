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

	var logger = helpers.Log(r)

	request, err := requests.NewPromocodeByIdRequest(r)
	if err != nil {
		logger.WithError(err).Error("failed to fetch rollback promocode request")
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
		logger.WithError(Inactive())
		ape.RenderErr(w, Inactive())
		return
	}
	promocodesQ := helpers.DB(r).Promocodes().New().UpdateUsages(promocode.Usages - 1)

	if promocode.Usages == 0 {
		promocodesQ = promocodesQ.UpdateState(resources.PromocodeActive)
	}

	if err = promocodesQ.FilterUpdateById(promocode.Id).Update(); err != nil {
		logger.WithError(err).Error("failed to get promocode")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
