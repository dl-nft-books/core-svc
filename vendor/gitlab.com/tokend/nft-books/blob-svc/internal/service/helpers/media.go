package helpers

import (
	"net/http"
	"strings"

	"gitlab.com/distributed_lab/logan/v3/errors"
)

func CheckDocumentMimeType(mtype string, r *http.Request) (string, error) {
	for _, el := range MimeTypes(r).AllowedMimeTypes {
		if el == mtype {
			return strings.Split(mtype, "/")[1], nil
		}
	}
	return "", errors.New("invalid file extension")
}
