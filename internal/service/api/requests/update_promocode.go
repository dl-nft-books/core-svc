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

	if _, err := strconv.Atoi(request.Data.ID); err != nil {
		return nil, errors.Wrap(err, "failed to get id from the url path")
	}

	return &request, request.validate()
}
func (r UpdatePromocodeRequest) validate() error {
	return validation.Errors{
		"data/attributes/discount": validation.Validate(
			r.Data.Attributes.Discount,
			validation.Min(0.0),
			validation.Max(1.0),
		),
	}.Filter()
}
