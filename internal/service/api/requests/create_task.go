package requests

import (
	"encoding/json"
	"github.com/dl-nft-books/core-svc/resources"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type CreateTaskRequest struct {
	resources.CreateTaskRequest
}

func NewCreateTaskRequest(r *http.Request) (*CreateTaskRequest, error) {
	var request CreateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal request")
	}

	return &request, request.validate()
}

func (r CreateTaskRequest) validate() error {
	return validation.Errors{
		"data/attributes/account": validation.Validate(
			&r.Data.Attributes.Account,
			validation.Required,
		),
		"data/attributes/book_id": validation.Validate(
			&r.Data.Attributes.BookId,
			validation.Required,
		),
		"data/attributes/chain_id": validation.Validate(
			&r.Data.Attributes.BookId,
			validation.Required,
		),
		"data/attributes/banner": validation.Validate(
			&r.Data.Attributes.BookId,
			validation.Required,
		),
	}.Filter()
}
