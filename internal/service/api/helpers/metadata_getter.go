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

func GetMetadataFromHash(r *http.Request, hash string) (*opensea.Metadata, error) {
	var (
		ipfser      = Ipfser(r)
		metadata    opensea.Metadata
		logger      = Log(r)
		loganFields = logan.F{
			"metadata_hash": hash,
		}
	)
	for i := 0; i < ipfser.NumberOfRetries; i++ {
		response, err := http.Get(ipfser.BaseUri + hash)
		if err != nil {
			if i+1 == ipfser.NumberOfRetries {
				return nil, errors.Wrap(err, "failed to get a response")
			}
			logger.WithFields(loganFields.Merge(logan.F{
				"attempts_number": 1,
			})).Error(errors.Wrap(err, "failed to get a response"))
			time.Sleep(ipfser.RetryPeriod)
			continue
		}

		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			if i+1 == ipfser.NumberOfRetries {
				return nil, errors.Wrap(err, "failed to read a response")
			}
			logger.WithFields(loganFields.Merge(logan.F{
				"attempts_number": 1,
			})).Error(errors.Wrap(err, "failed to read a response"))
			time.Sleep(ipfser.RetryPeriod)
			continue
		}

		if err = json.Unmarshal(responseData, &metadata); err != nil {
			if i+1 == ipfser.NumberOfRetries {
				return nil, errors.Wrap(err, "failed to unmarshal a response")
			}
			logger.WithFields(loganFields.Merge(logan.F{
				"attempts_number": 1,
			})).Error(errors.Wrap(err, "failed to unmarshal a response"))
			time.Sleep(ipfser.RetryPeriod)
			continue
		}
		break
	}

	return &metadata, nil
}
