package responses

import (
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	tracker "gitlab.com/tokend/nft-books/contract-tracker/connector"
	"gitlab.com/tokend/nft-books/generator-svc/internal/data"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/helpers"
	"gitlab.com/tokend/nft-books/generator-svc/resources"
	"net/http"
)

var (
	PaymentNotFoundErr = errors.New("payment with specified id was not found")
)

func NewGetTokenResponse(r *http.Request, token data.Token, trackerApi *tracker.Connector, beseUri string) (*resources.TokenResponse, error) {
	var response resources.TokenResponse

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

	metadata, err := helpers.GetMetadataFromHash(r, token.MetadataHash, beseUri)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get metadata from hash")
	}

	response.Data = resources.Token{
		Key: resources.NewKeyInt64(token.Id, resources.TOKENS),
		Attributes: resources.TokenAttributes{
			Owner:        token.Account,
			Description:  metadata.Description,
			MetadataHash: token.MetadataHash,
			ImageUrl:     metadata.Image,
			Name:         metadata.Name,
			Signature:    token.Signature,
			Status:       token.Status,
			TokenId:      token.TokenId,
		},
		Relationships: getTokenRelationships(token),
	}

	response.Included.Add(convertPaymentToResource(*paymentResponse))

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
