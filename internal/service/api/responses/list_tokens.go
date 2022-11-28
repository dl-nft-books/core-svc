package responses

import (
	"net/http"

	"gitlab.com/tokend/nft-books/generator-svc/internal/data/external"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/helpers"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/generator-svc/internal/data"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/requests"
	"gitlab.com/tokend/nft-books/generator-svc/resources"
)

func NewTokenListResponse(r *http.Request, request *requests.ListTokensRequest, tokens []data.Token, paymentsQ external.PaymentsQ, tasksQ data.TasksQ) (*resources.TokenListResponse, error) {
	response := resources.TokenListResponse{}

	if len(tokens) == 0 {
		return &resources.TokenListResponse{
			Data: make([]resources.Token, 0),
		}, nil
	}

	for _, token := range tokens {
		payment, err := paymentsQ.New().FilterById(token.PaymentId).Get()
		if err != nil {
			return nil, errors.Wrap(err, "failed to get payment by id", logan.F{
				"payment_id": token.PaymentId,
			})
		}
		if payment == nil {
			return nil, errors.From(PaymentNotFoundErr, logan.F{
				"payment_id": PaymentNotFoundErr,
			})
		}
		paymentAsResource := payment.Resource()

		tasks, err := tasksQ.New().Select(data.TaskSelector{
			IpfsHash: &token.MetadataHash,
		})
		if err != nil {
			return nil, errors.Wrap(err, "failed to get task by hash", logan.F{
				"metadata_hash": token.MetadataHash,
			})
		}
		if len(tasks) != 1 {
			return nil, errors.From(NonSingleTaskErr, logan.F{
				"metadata_hash": token.MetadataHash,
			})
		}
		task := tasks[0]

		metadata, err := helpers.GetMetadataFromHash(token.MetadataHash)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get metadata from hash")
		}

		tokenAsResource := resources.Token{
			Key: resources.NewKeyInt64(token.Id, resources.TOKENS),
			Attributes: resources.TokenAttributes{
				Owner:        token.Account,
				TokenId:      int32(token.TokenId),
				MetadataHash: token.MetadataHash,
				Status:       token.Status,
				Name:         metadata.Name,
				Description:  metadata.Description,
				ImageUrl:     metadata.Image,
				Signature:    task.Signature,
				ChainId:      token.ChainID,
			},
			Relationships: getTokenRelationships(token),
		}

		response.Data = append(response.Data, tokenAsResource)
		response.Included.Add(&paymentAsResource)
	}

	response.Links = requests.GetOffsetLinksWithSort(r, request.OffsetPageParams, request.Sorts)

	return &response, nil
}
