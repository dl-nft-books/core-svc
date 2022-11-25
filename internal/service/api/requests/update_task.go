package requests

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/generator-svc/resources"
	"net/http"
)

func NewUpdateTaskRequest(r *http.Request) (*resources.UpdateTaskRequest, error) {
	var request resources.UpdateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal update task request")
	}
	request.Data.ID = chi.URLParam(r, "id")

	// TODO: Add validation to request
	return &request, nil
}
