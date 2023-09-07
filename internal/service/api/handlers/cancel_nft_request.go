package handlers

import (
	"github.com/dl-nft-books/core-svc/resources"
	"net/http"

	"github.com/dl-nft-books/core-svc/internal/service/api/helpers"
	"github.com/dl-nft-books/core-svc/internal/service/api/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func CancelNftRequestById(w http.ResponseWriter, r *http.Request) {
	logger := helpers.Log(r)

	idToCancel, err := requests.NewNftRequestByIdRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	nftRequestsQ := helpers.DB(r).NftRequests()
	nftRequest, err := nftRequestsQ.FilterById(idToCancel).Get()
	if err != nil {
		logger.WithError(err).Error("failed to get nft request")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if nftRequest == nil {
		ape.RenderErr(w, problems.NotFound())
		return
	}
	if nftRequest.Status != resources.RequestPending {
		ape.RenderErr(w, problems.Forbidden())
		return
	}

	err = nftRequestsQ.UpdateStatus(resources.RequestCanceled).FilterUpdateById(idToCancel).Update()
	if err != nil {
		logger.WithError(err).Error("failed to update nft request")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
