package helpers

import (
	"encoding/json"
	"github.com/pkg/errors"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/runners/models"
	"io/ioutil"
	"net/http"
)

const baseURI = "https://ipfs.io/ipfs/"

func GetMetadataFromHash(hash string) (*models.Metadata, error) {
	response, err := http.Get(baseURI + hash)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get a response")
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read a response")
	}

	var metadata models.Metadata
	if err = json.Unmarshal(responseData, &metadata); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal a response")
	}

	return &metadata, nil
}
