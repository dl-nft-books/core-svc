package handlers

import (
	"net/http"

	"github.com/dl-nft-books/core-svc/internal/service/api/responses"

	"github.com/dl-nft-books/core-svc/internal/service/api/helpers"
	"github.com/dl-nft-books/core-svc/internal/service/api/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetNftRequestById(w http.ResponseWriter, r *http.Request) {
	logger := helpers.Log(r)

	idToGet, err := requests.NewNftRequestByIdRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	nftRequest, err := helpers.DB(r).NftRequests().FilterByMarketplaceId(idToGet).Get()
	if err != nil {
		logger.WithError(err).Error("failed to get nft request")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if nftRequest == nil {
		ape.RenderErr(w, problems.NotFound())
		return
	}

	nftRequestResponse := responses.NewGetNftRequestResponse(*nftRequest)

	ape.Render(w, *nftRequestResponse)
}
