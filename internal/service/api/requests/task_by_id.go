package requests

import (
	"github.com/go-chi/chi"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
)

type TaskByIdRequest struct {
	Id int64
}

func NewTaskByIdRequest(r *http.Request) (*TaskByIdRequest, error) {
	IdAsString, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		return nil, errors.Wrap(err, "failed to get id from the url path")
	}

	return &TaskByIdRequest{
		Id: int64(IdAsString),
	}, nil
}
