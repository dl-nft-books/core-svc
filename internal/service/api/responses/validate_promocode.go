package responses

import (
	"gitlab.com/tokend/nft-books/generator-svc/internal/data"
	"gitlab.com/tokend/nft-books/generator-svc/resources"
)

func NewValidatePromocodeResponse(promocode data.Promocode) (*resources.ValidatePromocodeResponse, error) {
	response := resources.ValidatePromocodeResponse{}

	response.Data = resources.ValidatePromocode{
		Key: resources.NewKeyInt64(promocode.Id, resources.VALIDATE_PROMOCODE),
		Attributes: resources.ValidatePromocodeAttributes{
			State: promocode.State,
		},
	}
	if promocode.State == resources.PromocodeActive {
		relKey := resources.NewKeyInt64(promocode.Id, resources.PROMOCODE)
		response.Data.Relationships = resources.ValidatePromocodeRelationships{
			Promocode: &resources.Relation{
				Data: &relKey,
			},
		}
		resPromocode := promocode.Resource()
		response.Included.Add(&resPromocode)
	}
	return &response, nil
}
