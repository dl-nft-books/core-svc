package requests

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"github.com/dl-nft-books/core-svc/resources"
)

func NewUpdateTaskRequest(r *http.Request) (*resources.UpdateTaskRequest, error) {
	var request resources.UpdateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal update task request")
	}

	request.Data.ID = chi.URLParam(r, "id")
	_, err := strconv.Atoi(request.Data.ID)
	if err != nil {
		return nil, errors.New("invalid id param")
	}

	return &request, nil
}
