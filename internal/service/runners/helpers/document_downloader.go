package helpers

import (
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"net/http"
)

func DownloadDocument(link string) (document []byte, err error) {
	request, err := http.NewRequest(http.MethodGet, link, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to form a download document request")
	}

	client := http.DefaultClient
	response, err := client.Do(request)
	if err != nil {
		return nil, errors.Wrap(err, "failed to perform get operation")
	}

	defer func(Body io.ReadCloser) {
		if tempErr := Body.Close(); tempErr != nil {
			err = tempErr
		}
	}(response.Body)

	return ioutil.ReadAll(response.Body)
}
