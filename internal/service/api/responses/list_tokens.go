package responses

import (
	"net/http"

	tracker "gitlab.com/tokend/nft-books/contract-tracker/connector"

	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/helpers"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/generator-svc/internal/data"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/requests"
	"gitlab.com/tokend/nft-books/generator-svc/resources"
)

func NewTokenListResponse(r *http.Request, request *requests.ListTokensRequest, tokens []data.Token, trackerApi *tracker.Connector) (*resources.TokenListResponse, error) {
	response := resources.TokenListResponse{}

	if len(tokens) == 0 {
		return &resources.TokenListResponse{
			Data: make([]resources.Token, 0),
		}, nil
	}

	for _, token := range tokens {
		paymentResponse, err := trackerApi.GetPaymentById(token.PaymentId)

		if err != nil {
			return nil, errors.Wrap(err, "failed to get payment by id", logan.F{
				"payment_id": token.PaymentId,
			})
		}
		if paymentResponse == nil {
			return nil, errors.From(PaymentNotFoundErr, logan.F{
				"payment_id": PaymentNotFoundErr,
			})
		}

		metadata, err := helpers.GetMetadataFromHash(r, token.MetadataHash, helpers.BaseIpfsUri(r))
		if err != nil {
			return nil, errors.Wrap(err, "failed to get metadata from hash", logan.F{
				"metadata_hash": token.MetadataHash,
			})
		}

		tokenAsResource := resources.Token{
			Key: resources.NewKeyInt64(token.Id, resources.TOKENS),
			Attributes: resources.TokenAttributes{
				Owner:        token.Account,
				TokenId:      token.TokenId,
				MetadataHash: token.MetadataHash,
				Status:       token.Status,
				Name:         metadata.Name,
				Description:  metadata.Description,
				ImageUrl:     metadata.Image,
				Signature:    token.Signature,
			},
			Relationships: getTokenRelationships(token),
		}

		response.Data = append(response.Data, tokenAsResource)
		response.Included.Add(convertPaymentToResource(*paymentResponse))
	}

	response.Links = requests.GetOffsetLinksWithSort(r, request.OffsetPageParams, request.Sorts)

	return &response, nil
}
