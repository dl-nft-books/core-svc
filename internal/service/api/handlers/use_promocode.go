package handlers

import (
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/tokend/nft-books/generator-svc/resources"
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/helpers"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/requests"
)

func UsePromocode(w http.ResponseWriter, r *http.Request) {
	logger := helpers.Log(r)

	request, err := requests.NewPromocodeByIdRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	promocode, err := helpers.DB(r).Promocodes().FilterById(request.Id).Get()
	if err != nil {
		logger.WithError(err).WithFields(logan.F{"promocode_id": request.Id}).Error("failed to get promocode")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if promocode == nil {
		ape.RenderErr(w, problems.NotFound())
		return
	}
	if promocode.State != resources.PromocodeActive {
		logger.WithError(err).WithFields(logan.F{"promocode": promocode.Promocode}).Info("promocode is inactive")
		errorInactive := problems.Forbidden()
		errorInactive.Detail = "promocode is inactive"
		ape.RenderErr(w, errorInactive)
		return
	}

	err = helpers.DB(r).Promocodes().New().UpdateLeftUsages(promocode.LeftUsages - 1).Update(promocode.Id)

	if err != nil {
		logger.WithError(err).WithFields(logan.F{"promocode": promocode.Promocode}).Error("failed to update promocode")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
