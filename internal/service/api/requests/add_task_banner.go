package requests

import (
	"github.com/go-chi/chi"
	"mime/multipart"
	"net/http"
	"strconv"

	"gitlab.com/distributed_lab/logan/v3/errors"
)

const DocumentKey = "Document"

func NewAddTaskBannerRequest(r *http.Request) (int64, multipart.File, error) {

	IdAsString, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		return 0, nil, errors.Wrap(err, "failed to get id from the url path")
	}
	id := int64(IdAsString)

	if err = r.ParseMultipartForm(1 << 32); err != nil {
		return 0, nil, errors.Wrap(err, "failed to parse document")
	}

	f, _, err := r.FormFile(DocumentKey)
	return id, f, err
}
