package helpers

import (
	"encoding/json"
	"github.com/pkg/errors"
	"gitlab.com/tokend/nft-books/generator-svc/internal/data/opensea"
	"io/ioutil"
	"net/http"
)

const baseURI = "https://ipfs.io/ipfs/"

func GetMetadataFromHash(hash string) (*opensea.Metadata, error) {
	response, err := http.Get(baseURI + hash)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get a response")
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read a response")
	}

	var metadata opensea.Metadata
	if err = json.Unmarshal(responseData, &metadata); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal a response")
	}

	return &metadata, nil
}
