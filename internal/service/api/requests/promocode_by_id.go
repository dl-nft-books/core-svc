package requests

import (
	"github.com/go-chi/chi"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
)

type PromocodeByIdRequest struct {
	Id int64
}

func NewPromocodeByIdRequest(r *http.Request) (*PromocodeByIdRequest, error) {
	IdAsString, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		return nil, errors.Wrap(err, "failed to get id from the url path")
	}

	return &PromocodeByIdRequest{
		Id: int64(IdAsString),
	}, nil
}
