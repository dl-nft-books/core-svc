package requests

import (
	"encoding/json"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/generator-svc/resources"
)

type UpdatePromocodeRequest struct {
	resources.UpdatePromocodeRequest
}

func NewUpdatePromocodeRequest(r *http.Request) (*UpdatePromocodeRequest, error) {
	var request UpdatePromocodeRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal update task request")
	}

	request.Data.ID = chi.URLParam(r, "id")
	_, err := strconv.Atoi(request.Data.ID)
	if err != nil {
		return nil, errors.New("invalid id param")
	}

	return &request, request.validate()
}
func (r UpdatePromocodeRequest) validate() error {
	return validation.Errors{
		"data/attributes/discount": validation.Validate(
			r.Data.Attributes.Discount,
			validation.Required,
			validation.Min(0.0),
			validation.Max(1.0),
		),
		"data/attributes/initial_usages": validation.Validate(
			r.Data.Attributes.InitialUsages,
			validation.Required,
		),
		"data/attributes/expiration_date": validation.Validate(
			r.Data.Attributes.ExpirationDate,
			validation.Required,
		),
	}.Filter()
}
