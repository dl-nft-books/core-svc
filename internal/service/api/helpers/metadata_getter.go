package helpers

import (
	"encoding/json"
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/tokend/nft-books/generator-svc/internal/data/opensea"
	"io/ioutil"
	"net/http"
	"time"
)

func GetMetadataFromHash(r *http.Request, hash, baseUri string) (*opensea.Metadata, error) {
	var (
		numberOfRetries = 5
		retryAfter      = 2 * time.Second
		metadata        opensea.Metadata
		logger          = Log(r)
	)
	for i := 0; i < numberOfRetries; i++ {
		response, err := http.Get(baseUri + hash)
		if err != nil {
			if i+1 == numberOfRetries {
				return nil, errors.Wrap(err, "failed to get a response")
			}
			logger.WithFields(logan.F{
				"try number":    i + 1,
				"metadata_hash": hash,
			}).Error(errors.Wrap(err, "failed to get a response"))
			continue
		}

		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			if i+1 == numberOfRetries {
				return nil, errors.Wrap(err, "failed to read a response")
			}
			logger.WithFields(logan.F{
				"try number":    i + 1,
				"metadata_hash": hash,
			}).Error(errors.Wrap(err, "failed to read a response"))
			continue
		}

		if err = json.Unmarshal(responseData, &metadata); err != nil {
			if i+1 == numberOfRetries {
				return nil, errors.Wrap(err, "failed to unmarshal a response")
			}
			logger.WithFields(logan.F{
				"try number":    i + 1,
				"metadata_hash": hash,
			}).Error(errors.Wrap(err, "failed to unmarshal a response"))
			continue
		}
		if err == nil {
			break
		}
		time.Sleep(retryAfter)
	}

	return &metadata, nil
}
