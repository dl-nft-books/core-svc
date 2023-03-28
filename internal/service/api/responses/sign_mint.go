package responses

import (
	"github.com/dl-nft-books/core-svc/internal/signature"
	"github.com/dl-nft-books/core-svc/resources"
)

func NewSignMintResponse(price string, discount string, signature *signature.Parameters, endTimestamp int64) resources.PriceResponse {
	priceKey := resources.NewKeyInt64(1, resources.PRICES)
	signatureKey := resources.NewKeyInt64(1, resources.SIGNATURES)

	included := resources.Included{}
	included.Add(&resources.Signature{
		Key: signatureKey,
		Attributes: resources.SignatureAttributes{
			R: signature.R,
			S: signature.S,
			V: signature.V,
		},
	})

	return resources.PriceResponse{
		Data: resources.Price{
			Key: priceKey,
			Attributes: resources.PriceAttributes{
				Price:        price,
				EndTimestamp: endTimestamp,
				Discount:     discount,
			},
			Relationships: resources.PriceRelationships{
				Signature: resources.Relation{
					Data: &signatureKey,
				},
			},
		},
		Included: included,
	}
}
