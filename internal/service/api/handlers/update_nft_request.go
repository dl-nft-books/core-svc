package handlers

import (
	"github.com/spf13/cast"
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
		logger.WithError(err).Error("failed to fetch update nft request request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	nftRequestId := cast.ToInt64(request.Data.ID)
	nftRequest, err := helpers.DB(r).NftRequests().FilterById(nftRequestId).Get()
	if err != nil {
		logger.WithError(err).Error("failed to get nft request")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if nftRequest == nil {
		ape.RenderErr(w, problems.NotFound())
		return
	}

	if err = helpers.DB(r).NftRequests().New().
		UpdateStatus(request.Data.Attributes.Status).
		FilterUpdateById(nftRequestId).
		Update(); err != nil {
		logger.WithError(err).Error("failed to update nft request")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
