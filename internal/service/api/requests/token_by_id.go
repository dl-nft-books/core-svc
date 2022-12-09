package requests

import (
	"github.com/go-chi/chi"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
)

type TokenByIdRequest struct {
	Id int64
}

func NewTokenByIdRequest(r *http.Request) (*TokenByIdRequest, error) {
	IdAsString, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		return nil, errors.Wrap(err, "failed to get id from the url path")
	}

	return &TokenByIdRequest{
		Id: int64(IdAsString),
	}, nil
}
