package responses

import (
	"net/http"

	"github.com/dl-nft-books/core-svc/internal/data"
	"github.com/dl-nft-books/core-svc/internal/service/api/requests"
	"github.com/dl-nft-books/core-svc/resources"
)

func NewNftRequestListResponse(r *http.Request, request *requests.ListNftRequestsRequest, nftRequests []data.NftRequest) (*resources.NftRequestListResponse, error) {
	response := resources.NftRequestListResponse{}

	if len(nftRequests) == 0 {
		return &resources.NftRequestListResponse{
			Data: make([]resources.NftRequest, 0),
		}, nil
	}

	for _, nftRequest := range nftRequests {
		response.Data = append(response.Data, nftRequest.Resource())
	}

	response.Links = requests.GetOffsetLinksWithSort(r, request.OffsetPageParams, request.Sorts)

	return &response, nil
}
