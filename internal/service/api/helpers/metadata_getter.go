package helpers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/pkg/errors"
	"gitlab.com/tokend/nft-books/generator-svc/internal/data/opensea"
)

func GetMetadataFromHash(hash, baseUri string) (*opensea.Metadata, error) {
	log.Println("GET METADATA", baseUri+hash)
	response, err := http.Get(baseUri + hash)
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
