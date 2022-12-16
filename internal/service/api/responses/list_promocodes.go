package responses

import (
	"net/http"

	"gitlab.com/tokend/nft-books/generator-svc/internal/data"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/requests"
	"gitlab.com/tokend/nft-books/generator-svc/resources"
)

func NewPromocodeListResponse(r *http.Request, request *requests.ListPromocodesRequest, promocodes []data.Promocode) (*resources.PromocodeListResponse, error) {
	response := resources.PromocodeListResponse{}

	if len(promocodes) == 0 {
		return &resources.PromocodeListResponse{
			Data: make([]resources.Promocode, 0),
		}, nil
	}

	for _, promocode := range promocodes {
		promocodeAsResource := promocode.Resource()
		response.Data = append(response.Data, promocodeAsResource)
	}

	response.Links = requests.GetOffsetLinksWithSort(r, request.OffsetPageParams, request.Sorts)

	return &response, nil
}
