package handlers

import (
	"github.com/dl-nft-books/core-svc/internal/data"
	"github.com/dl-nft-books/core-svc/internal/service/api/helpers"
	"github.com/dl-nft-books/core-svc/internal/service/api/requests"
	"github.com/dl-nft-books/core-svc/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"net/http"
	"strconv"
	"time"
)

const recreatingRequestInterval = time.Hour * 24

func CreateNftRequest(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewCreateNftRequestRequest(r)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to fetch create nft request request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	nftReq, err := helpers.DB(r).NftRequests().
		FilterByPayerAddress(request.Data.Attributes.PayerAddress).
		FilterByCollectionAddress(request.Data.Attributes.CollectionAddress).
		FilterByNftId(request.Data.Attributes.NftId).Get()

	if nftReq != nil && nftReq.Status != resources.RequestRejected {
		helpers.Log(r).WithError(err).Error("request with such nft is already exists")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	if nftReq != nil && time.Now().Sub(nftReq.LastUpdatedAt) < recreatingRequestInterval {
		helpers.Log(r).WithError(err).Error("request with such nft is already exists")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	bookId, err := strconv.Atoi(request.Data.Relationships.Book.Data.ID)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to convert book_id")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	nftRequestID, err := helpers.DB(r).NftRequests().Insert(data.NftRequest{
		PayerAddress:      request.Data.Attributes.PayerAddress,
		CollectionAddress: request.Data.Attributes.CollectionAddress,
		NftId:             request.Data.Attributes.NftId,
		ChainId:           request.Data.Attributes.ChainId,
		BookId:            int64(bookId),
		Status:            resources.RequestPending,
		CreatedAt:         time.Now(),
		LastUpdatedAt:     time.Now(),
	})
	if err != nil {
		helpers.Log(r).WithError(err).Errorf("failed to create new nft request")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, resources.KeyResponse{
		Data: resources.NewKeyInt64(nftRequestID, resources.NFT_REQUEST),
	})
}
