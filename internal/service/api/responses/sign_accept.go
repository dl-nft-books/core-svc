package responses

import (
	"github.com/dl-nft-books/core-svc/internal/signature"
	"github.com/dl-nft-books/core-svc/resources"
)

func NewSignAcceptResponse(signature *signature.Parameters) resources.SignatureResponse {
	return resources.SignatureResponse{
		Data: resources.Signature{
			Key: resources.Key{
				Type: resources.SIGNATURES,
			},
			Attributes: resources.SignatureAttributes{
				R: signature.R,
				S: signature.S,
				V: signature.V,
			},
		},
	}
}
