package requests

import (
	"encoding/json"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/generator-svc/resources"
)

type CreatePromocodeRequest struct {
	resources.CreatePromocodeRequest
}

func NewCreatePromocodeRequest(r *http.Request) (*CreatePromocodeRequest, error) {
	var request CreatePromocodeRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal create promocode request")
	}

	return &request, request.validate()
}

func (r CreatePromocodeRequest) validate() error {
	return validation.Errors{
		"data/attributes/discount": validation.Validate(
			&r.Data.Attributes.Discount,
			validation.Required,
			validation.Min(0.0),
			validation.Max(1.0),
		),
		"data/attributes/initial_usages": validation.Validate(
			&r.Data.Attributes.InitialUsages,
			validation.Required,
		),
		"data/attributes/expiration_date": validation.Validate(
			&r.Data.Attributes.ExpirationDate,
			validation.Required,
		),
	}.Filter()
}
