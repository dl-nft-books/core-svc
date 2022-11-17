package responses

import (
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/generator-svc/internal/data"
	"gitlab.com/tokend/nft-books/generator-svc/internal/data/external"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/helpers"
	"gitlab.com/tokend/nft-books/generator-svc/resources"
)

const baseUri = "https://ipfs.io/ipfs/"

var MultipleOrNoneTasksErr = errors.New("Either no tasks or duplicate for the given hash were found")

func NewGetTokenResponse(token data.Token, paymentsQ external.PaymentsQ, tasksQ data.TasksQ) (*resources.TokenResponse, error) {
	var response resources.TokenResponse

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

	response.Data = resources.Token{
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

	response.Included.Add(paymentAsResource)

	return &response, nil
}

func getTokenRelationships(token data.Token) resources.TokenRelationships {
	var (
		bookKey    = resources.NewKeyInt64(token.BookId, resources.BOOKS)
		paymentKey = resources.NewKeyInt64(token.PaymentId, resources.PAYMENT)
	)

	return resources.TokenRelationships{
		Book: &resources.Relation{
			Data: &bookKey,
		},
		Payment: &resources.Relation{
			Data: &paymentKey,
		},
	}
}
