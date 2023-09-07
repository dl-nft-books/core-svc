package handlers

import (
	"github.com/dl-nft-books/core-svc/internal/data"
	"github.com/dl-nft-books/core-svc/internal/service/api/helpers"
	"gitlab.com/distributed_lab/ape/problems"
	"net/http"

	"github.com/dl-nft-books/core-svc/internal/service/api/requests"
	"github.com/dl-nft-books/core-svc/internal/service/api/responses"
	"gitlab.com/distributed_lab/ape"
)

func ListNftRequests(w http.ResponseWriter, r *http.Request) {
	logger := helpers.Log(r)

	request, err := requests.NewListNftRequestsRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	nftRequests, err := applyNftRequestsQFilters(helpers.DB(r).NftRequests(), request).Select()
	if err != nil {
		logger.WithError(err).Error("unable to select nft requests from database")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	nftRequestsListResponse, err := responses.NewNftRequestListResponse(r, request, nftRequests, *helpers.Booker(r))
	if err != nil {
		logger.WithError(err).Error("unable to form nftRequests list response")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	ape.Render(w, *nftRequestsListResponse)
}

func applyNftRequestsQFilters(q data.NftRequestsQ, request *requests.ListNftRequestsRequest) data.NftRequestsQ {
	if len(request.Status) != 0 {
		q = q.FilterByStatus(request.Status...)
	}
	if len(request.BookId) != 0 {
		q = q.FilterById(request.BookId...)
	}
	if len(request.ChainId) != 0 {
		q = q.FilterByChainId(request.ChainId...)
	}
	if len(request.Requester) != 0 {
		q = q.FilterByRequester(request.Requester...)
	}
	if len(request.NftAddress) != 0 {
		q = q.FilterByNftAddress(request.NftAddress...)
	}

	q = q.Page(request.OffsetPageParams)
	q = q.Sort(request.Sorts)

	return q
}
