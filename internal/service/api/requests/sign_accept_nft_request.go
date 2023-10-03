package requests

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/urlval"
	"net/http"
)

type SignAcceptNftRequestRequest struct {
	TaskID       int64 `url:"task_id"`
	NftRequestID int64 `url:"request_id"`
}

func NewSignAcceptNftRequestRequest(r *http.Request) (*SignAcceptNftRequestRequest, error) {
	var result SignAcceptNftRequestRequest
	var err error

	if err = urlval.Decode(r.URL.Query(), &result); err != nil {
		return &result, errors.Wrap(err, "failed to unmarshal sign accept request")
	}

	return &result, result.validate()
}

func (r SignAcceptNftRequestRequest) validate() error {
	return validation.Errors{
		"task_id=": validation.Validate(
			r.TaskID,
			validation.Required,
			validation.Min(1)),
		"request_id=": validation.Validate(
			r.NftRequestID,
			validation.Required,
			validation.Min(1)),
	}.Filter()

}
