package requests

import (
	"net/http"

	"github.com/go-chi/chi"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/urlval"
)

type GetPriceRequest struct {
	ID int64 `json:"id"`

	ChainID      int64  `json:"chain_id"`
	TokenAddress string `json:"token_address"`
}

func NewGetPriceRequest(r *http.Request) (*GetPriceRequest, error) {
	var result GetPriceRequest

	err := urlval.Decode(r.URL.Query(), &result)
	if err != nil {
		return nil, err
	}

	result.ID, err = cast.ToInt64E(chi.URLParam(r, "id"))
	if err != nil {
		return nil, validation.Errors{
			"id": err,
		}
	}

	return &result, nil
}
