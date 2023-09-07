package handlers

import (
	"errors"
	"github.com/dl-nft-books/core-svc/internal/data"
	"github.com/dl-nft-books/core-svc/internal/service/api/jsonerrors"
	"github.com/dl-nft-books/core-svc/resources"
	"gitlab.com/distributed_lab/logan/v3"
	"net/http"

	"github.com/dl-nft-books/core-svc/internal/service/api/helpers"
	"github.com/dl-nft-books/core-svc/internal/service/api/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func UpdateNftRequestById(w http.ResponseWriter, r *http.Request) {
	logger := helpers.Log(r)

	request, err := requests.NewUpdateNftRequestRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	statusToSet := request.Data.Attributes.Status
	if statusToSet == resources.RequestAccepted || statusToSet == resources.RequestRejected {
		address := r.Context().Value("address").(string)
		isMarketplaceManager, err := helpers.CheckMarketplacePerrmision(*helpers.Networker(r), address)
		if err != nil {
			helpers.Log(r).WithError(err).WithFields(logan.F{"account": address}).Debug("failed to check permissions")
			ape.RenderErr(w, problems.InternalError())
			return
		}
		if !isMarketplaceManager {
			helpers.Log(r).WithFields(logan.F{"account": address}).Debug("you don't have permission to set the status")
			ape.RenderErr(w, jsonerrors.WithDetails(problems.Forbidden(), jsonerrors.NotManagerAuthToken))
			return
		}
	}

	nftRequestsQ := helpers.DB(r).NftRequests().New()
	if err = nftRequestsQ.UpdateStatus(request.Data.Attributes.Status).FilterUpdateById(request.ID).Update(); err != nil {
		if errors.Is(err, data.NoRowsAffected) {
			ape.RenderErr(w, problems.NotFound())
			return
		}
		logger.WithError(err).Error("failed to update nft request")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
