package responses

import (
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/helpers"
	"net/http"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/generator-svc/internal/data"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/requests"
	"gitlab.com/tokend/nft-books/generator-svc/resources"
)

func NewTokenListResponse(r *http.Request, request *requests.ListTokensRequest, tokens []data.Token, paymentsQ data.PaymentsQ, tasksQ data.TasksQ) (*resources.TokenListResponse, error) {
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
		paymentAsResource, err := payment.Resource()
		if err != nil {
			return nil, errors.Wrap(err, "failed to convert payment to the resource format")
		}

		tasks, err := tasksQ.New().Select(data.TaskSelector{
			IpfsHash: &token.MetadataHash,
		})
		if err != nil {
			return nil, errors.Wrap(err, "failed to get task by hash", logan.F{
				"metadata_hash": token.MetadataHash,
			})
		}
		if len(tasks) != 1 {
			return nil, errors.From(MultipleOrNoneTasksErr, logan.F{
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
				Description: metadata.Description,
				ImageUrl:    metadata.Image,
				Name:        metadata.Name,
				Signature:   task.Signature,
				Status:      token.Status,
				TokenId:     int32(token.TokenId),
			},
			Relationships: getTokenRelationships(token),
		}

		response.Data = append(response.Data, tokenAsResource)
		response.Included.Add(paymentAsResource)
	}

	response.Links = requests.GetOffsetLinksWithSort(r, request.OffsetPageParams, request.Sorts)

	return &response, nil
}
