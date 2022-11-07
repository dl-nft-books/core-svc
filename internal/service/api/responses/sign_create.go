package responses

import (
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/signature"
	"gitlab.com/tokend/nft-books/generator-svc/resources"
)

func NewSignCreateResponse(tokenId int64, signature *signature.SignatureParameters) resources.CreateSignatureResponse {
	createSignatureKey := resources.NewKeyInt64(1, resources.CREATE_SIGNATURES)
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

	return resources.CreateSignatureResponse{
		Data: resources.CreateSignature{
			Key: createSignatureKey,
			Attributes: resources.CreateSignatureAttributes{
				TokenId: int32(tokenId),
			},
			Relationships: resources.CreateSignatureRelationships{
				Signature: resources.Relation{
					Data: &signatureKey,
				},
			},
		},
		Included: included,
	}
}
