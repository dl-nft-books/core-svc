package requests

import (
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/urlval"
)

type SignCreateRequest struct {
	TokenName        string `url:"token_name"`
	TokenSymbol      string `url:"token_symbol"`
	PricePerOneToken string `url:"price"`
}

func NewSignCreateRequest(r *http.Request) (*SignCreateRequest, error) {
	var result SignCreateRequest
	var err error

	if err = urlval.Decode(r.URL.Query(), &result); err != nil {
		return &result, err
	}

	return &result, result.validate()
}

func (r SignCreateRequest) validate() error {
	return validation.Errors{
		"token_name=":   validation.Validate(r.TokenName, validation.Required),
		"token_symbol=": validation.Validate(r.TokenSymbol, validation.Required),
		"price=":        validation.Validate(r.PricePerOneToken, validation.Required),
	}.Filter()
}
