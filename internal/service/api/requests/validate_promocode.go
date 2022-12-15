package requests

import (
	"github.com/go-chi/chi"
	"net/http"
)

type ValidatePromocodeRequest struct {
	Promocode string
}

func NewValidatePromocodeRequest(r *http.Request) (*ValidatePromocodeRequest, error) {
	promocode := chi.URLParam(r, "promocode")

	return &ValidatePromocodeRequest{
		Promocode: promocode,
	}, nil
}
