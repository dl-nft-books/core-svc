package requests

import (
	"encoding/json"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/generator-svc/resources"
)

type CreateTokenRequest struct {
	resources.CreateTokenRequest
}

func NewCreateTokenRequest(r *http.Request) (*CreateTokenRequest, error) {
	var request CreateTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal create token request")
	}

	return &request, request.validate()
}

func (r CreateTokenRequest) validate() error {
	return validation.Errors{
		"data/attributes/token_id": validation.Validate(
			&r.Data.Attributes.TokenId,
			validation.Required,
		),
		"data/attributes/status": validation.Validate(
			&r.Data.Attributes.Status,
			validation.Required,
		),
	}.Filter()
}
