package requests

import (
	"net/http"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/urlval"
)

var AddressRegexp = regexp.MustCompile("^(0x)?[0-9a-fA-F]{40}$")

type SignMintRequest struct {
	TaskID       int64  `url:"task_id"`
	Platform     string `url:"platform"`
	TokenAddress string `url:"token_address"`
	PromocodeID  int64  `url:"promocode_id"`
}

func NewSignMintRequest(r *http.Request) (*SignMintRequest, error) {
	var result SignMintRequest
	var err error

	if err = urlval.Decode(r.URL.Query(), &result); err != nil {
		return &result, err
	}

	return &result, result.validate()
}

func (r SignMintRequest) validate() error {
	err := validation.Errors{
		"task_id=": validation.Validate(
			r.TaskID,
			validation.Required,
			validation.Min(1)),
		"platform=": validation.Validate(r.Platform, validation.Required),
	}.Filter()

	if err != nil {
		return err
	}

	if r.TokenAddress != "" {
		return validation.Errors{
			"token_address=": validation.Validate(
				r.TokenAddress,
				validation.Required,
				validation.Match(AddressRegexp)),
		}.Filter()
	}

	return nil
}
