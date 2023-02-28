package requests

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/generator-svc/resources"
)

func NewUpdateTokenRequest(r *http.Request) (*resources.UpdateTokenRequest, error) {
	var request resources.UpdateTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal update token request")
	}

	request.Data.ID = chi.URLParam(r, "id")
	_, err := strconv.Atoi(request.Data.ID)
	if err != nil {
		return nil, errors.New("invalid id param")
	}

	return &request, nil
}
