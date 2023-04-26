package handlers

import (
	"fmt"
	"github.com/dl-nft-books/core-svc/resources"
	"github.com/spf13/cast"
	"net/http"

	"github.com/dl-nft-books/core-svc/internal/service/api/helpers"
	"github.com/dl-nft-books/core-svc/internal/service/api/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func CancelNftRequestById(w http.ResponseWriter, r *http.Request) {
	logger := helpers.Log(r)

	request, err := requests.NewNftRequestByIdRequest(r)
	if err != nil {
		logger.WithError(err).Error("failed to fetch update nft request request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	nftRequestId := cast.ToInt64(request.Id)
	nftRequest, err := helpers.DB(r).NftRequests().FilterById(nftRequestId).Get()
	if err != nil {
		logger.WithError(err).Error("failed to get nft request")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if nftRequest == nil {
		logger.WithError(err).Error("nft request with such id not found")
		ape.RenderErr(w, problems.NotFound())
		return
	}
	if nftRequest.Status != resources.RequestPending {
		logger.WithError(err).Error(fmt.Sprintf("can not canceled nft request with status %v", nftRequest.Status))
		ape.RenderErr(w, problems.Forbidden())
		return
	}

	if err = helpers.DB(r).NftRequests().New().
		UpdateStatus(resources.RequestCanceled).
		FilterUpdateById(nftRequestId).
		Update(); err != nil {
		logger.WithError(err).Error("failed to update nft request")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
