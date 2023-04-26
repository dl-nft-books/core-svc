package handlers

import (
	"gitlab.com/distributed_lab/logan/v3"
	"net/http"

	"github.com/dl-nft-books/core-svc/internal/service/api/helpers"
	"github.com/dl-nft-books/core-svc/internal/service/api/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func DeletePromocodeById(w http.ResponseWriter, r *http.Request) {
	logger := helpers.Log(r)

	request, err := requests.NewPromocodeByIdRequest(r)
	if err != nil {
		logger.WithError(err).Error("failed to fetch delete promocode request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	address := r.Context().Value("address").(string)
	isMarketplaceManager, err := helpers.CheckMarketplacePerrmision(*helpers.Networker(r), address)
	if err != nil {
		helpers.Log(r).WithError(err).WithFields(logan.F{"account": address}).Debug("failed to check permissions")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if !isMarketplaceManager {
		helpers.Log(r).WithFields(logan.F{"account": address}).Debug("you don't have permission to create book")
		ape.RenderErr(w, problems.Forbidden())
		return
	}
	promocode, err := helpers.DB(r).Promocodes().FilterById(request.Id).Get()
	if err != nil {
		logger.WithError(err).Error("failed to get promocode")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if promocode == nil {
		logger.WithFields(logan.F{"promocode_id": request.Id}).Error("promocode with such id is not found")
		ape.RenderErr(w, problems.NotFound())
		return
	}

	if err = helpers.DB(r).Promocodes().DeleteByID(request.Id); err != nil {
		logger.WithError(err).Error("failed to delete promocode")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
