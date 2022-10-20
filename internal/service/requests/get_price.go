package requests

import (
	"net/http"
	"regexp"

	"github.com/go-chi/chi"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

var AddressRegexp = regexp.MustCompile("^(0x)?[0-9a-fA-F]{40}$")

type GetPriceRequest struct {
	ID           int64  `json:"id"`
	Platform     string `json:"platform"`
	TokenAddress string `json:"token_address"`
}

func NewGetPriceRequest(r *http.Request) (*GetPriceRequest, error) {
	var result GetPriceRequest
	var err error

	result.ID, err = cast.ToInt64E(chi.URLParam(r, "id"))
	if err != nil {
		return nil, validation.Errors{
			"id": err,
		}
	}

	result.Platform = r.URL.Query().Get("platform")
	if result.Platform == "" {
		return nil, errors.New("failed to retrieve platform parameter")
	}

	result.TokenAddress = r.URL.Query().Get("token_address")

	return &result, result.validate()
}

func (r GetPriceRequest) validate() error {
	if r.TokenAddress == "" {
		return nil
	}

	return validation.Errors{
		"token_address": validation.Validate(
			&r.TokenAddress,
			validation.Required,
			validation.Match(AddressRegexp)),
	}.Filter()
}
