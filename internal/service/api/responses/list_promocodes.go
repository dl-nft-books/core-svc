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
		promocodeAsResource := resources.Promocode{
			Key: resources.NewKeyInt64(promocode.Id, resources.PROMOCODE),
			Attributes: resources.PromocodeAttributes{
				Id:             promocode.Id,
				Promocode:      promocode.Promocode,
				Discount:       promocode.Discount,
				InitialUsages:  promocode.InitialUsages,
				LeftUsages:     &promocode.LeftUsages,
				ExpirationDate: promocode.ExpirationDate,
				State:          &promocode.State,
			},
		}

		response.Data = append(response.Data, promocodeAsResource)
	}

	response.Links = requests.GetOffsetLinksWithSort(r, request.OffsetPageParams, request.Sorts)

	return &response, nil
}
