package responses

import (
	"gitlab.com/tokend/nft-books/generator-svc/internal/data"
	"gitlab.com/tokend/nft-books/generator-svc/resources"
)

func NewGetPromocodeResponse(promocode data.Promocode) (*resources.PromocodeResponse, error) {
	response := resources.PromocodeResponse{}

	response.Data = resources.Promocode{
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

	return &response, nil
}
