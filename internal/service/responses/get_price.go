package responses

import (
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/helpers"
	"gitlab.com/tokend/nft-books/generator-svc/resources"
)

func NewGetPriceResponse(tokenAmount string, signature *helpers.SignatureParameters) resources.PriceResponse {
	priceKey := resources.NewKeyInt64(1, resources.PRICES)
	signatureKey := resources.NewKeyInt64(1, resources.SIGNATURES)

	included := resources.Included{}
	included.Add(&resources.Signature{
		Key: signatureKey,
		Attributes: resources.SignatureAttributes{
			R: signature.R,
			S: signature.S,
			V: int32(signature.V),
		},
	})

	return resources.PriceResponse{
		Data: resources.Price{
			Key: priceKey,
			Attributes: resources.PriceAttributes{
				Price: tokenAmount,
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
