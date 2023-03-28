package requests

import (
	"github.com/go-chi/chi"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
)

type NftRequestByIdRequest struct {
	Id int64
}

func NewNftRequestByIdRequest(r *http.Request) (*NftRequestByIdRequest, error) {
	idAsString, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		return nil, errors.Wrap(err, "failed to get id from the url path")
	}

	return &NftRequestByIdRequest{
		Id: int64(idAsString),
	}, nil
}
