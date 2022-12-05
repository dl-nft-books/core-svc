package requests

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type GetPaymentRequest struct {
	Id int64 `json:"-"`
}

func NewGetPaymentRequest(r *http.Request) (*GetPaymentRequest, error) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		return nil, validation.Errors{"id": errors.New("id should be a valid int64")}
	}

	return &GetPaymentRequest{Id: id}, nil
}
