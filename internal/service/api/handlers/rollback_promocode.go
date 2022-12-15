package handlers

import (
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/responses"
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
		ape.RenderErr(w, problems.Forbidden())
		return
	}

	promocodesQ := helpers.DB(r).Promocodes().New().UpdateLeftUsages(promocode.LeftUsages + 1)

	if promocode.LeftUsages == 0 {
		promocodesQ = promocodesQ.UpdateState(resources.PromocodeActive)
	}

	promocode, err = promocodesQ.Update(promocode.Id)

	if err != nil {
		logger.WithError(err).Error("failed to get promocode")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if promocode == nil {
		ape.RenderErr(w, problems.NotFound())
		return
	}
	promocodeResponse, err := responses.NewGetPromocodeResponse(*promocode)
	if err != nil {
		logger.WithError(err).Error("failed to get promocode response")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, *promocodeResponse)
}
