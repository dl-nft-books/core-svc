package helpers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"gitlab.com/tokend/nft-books/generator-svc/internal/data/opensea"
)

func GetMetadataFromHash(hash, baseUri string) (*opensea.Metadata, error) {
	log.Println("GET METADATA", baseUri+hash)
	var metadata opensea.Metadata
	for {
		response, err := http.Get(baseUri + hash)
		if err != nil {
			//return nil, errors.Wrap(err, "failed to get a response")
			continue
		}

		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			//return nil, errors.Wrap(err, "failed to read a response")
			continue
		}

		if err = json.Unmarshal(responseData, &metadata); err != nil {
			//return nil, errors.Wrap(err, "failed to unmarshal a response")
			continue
		}
		if err == nil {
			break
		}
		time.Sleep(2 * time.Second)
	}

	return &metadata, nil
}
