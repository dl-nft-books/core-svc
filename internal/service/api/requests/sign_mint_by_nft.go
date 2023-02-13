package requests

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/urlval"
	"math/big"
	"net/http"
)

type SignMintByNftRequest struct {
	TaskID     int64    `url:"task_id"`
	Platform   string   `url:"platform"`
	NftAddress string   `url:"token_address"`
	NftID      *big.Int `url:"nft_id"`
}

func NewSignMintByNftRequest(r *http.Request) (*SignMintByNftRequest, error) {
	var result SignMintByNftRequest
	var err error

	if err = urlval.Decode(r.URL.Query(), &result); err != nil {
		return &result, err
	}

	return &result, result.validate()
}

func (r SignMintByNftRequest) validate() error {
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

	return nil
}
