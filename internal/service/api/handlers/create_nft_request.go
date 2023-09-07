package handlers

import (
	"github.com/dl-nft-books/core-svc/internal/data"
	"github.com/dl-nft-books/core-svc/internal/service/api/helpers"
	"github.com/dl-nft-books/core-svc/internal/service/api/requests"
	"github.com/dl-nft-books/core-svc/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"net/http"
	"time"
)

func CreateNftRequest(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewCreateNftRequestRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	nftRequestID, err := helpers.DB(r).NftRequests().Insert(data.NftRequest{
		Requester:            request.Data.Attributes.Requester,
		MarketplaceRequestId: request.Data.Attributes.MarketplaceRequestId,
		NftAddress:           request.Data.Attributes.NftAddress,
		NftId:                request.Data.Attributes.NftId,
		ChainId:              request.Data.Attributes.ChainId,
		BookId:               request.Data.Attributes.BookId,
		Status:               resources.RequestPending,
		CreatedAt:            time.Now(),
		LastUpdatedAt:        time.Now(),
	})
	if err != nil {
		helpers.Log(r).WithError(err).Errorf("failed to create new nft request")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	w.WriteHeader(http.StatusCreated)
	ape.Render(w, resources.KeyResponse{
		Data: resources.NewKeyInt64(nftRequestID, resources.NFT_REQUEST),
	})
}
