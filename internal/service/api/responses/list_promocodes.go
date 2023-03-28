package responses

import (
	"net/http"

	"github.com/dl-nft-books/core-svc/internal/data"
	"github.com/dl-nft-books/core-svc/internal/service/api/requests"
	"github.com/dl-nft-books/core-svc/resources"
)

func NewPromocodeListResponse(r *http.Request, request *requests.ListPromocodesRequest, promocodes []data.Promocode) (*resources.PromocodeListResponse, error) {
	response := resources.PromocodeListResponse{}

	if len(promocodes) == 0 {
		return &resources.PromocodeListResponse{
			Data: make([]resources.Promocode, 0),
		}, nil
	}

	for _, promocode := range promocodes {
		response.Data = append(response.Data, promocode.Resource())
	}

	response.Links = requests.GetOffsetLinksWithSort(r, request.OffsetPageParams, request.Sorts)

	return &response, nil
}
