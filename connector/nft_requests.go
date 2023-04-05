package connector

import (
	"encoding/json"
	"fmt"
	"github.com/dl-nft-books/core-svc/connector/models"
	"github.com/dl-nft-books/core-svc/resources"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/urlval"
)

const (
	nftRequestEndpoint = "nft-request"
)

func (c *Connector) UpdateNFTRequestStatus(id int64, status resources.NftRequestStatus) error {
	request := resources.UpdateNftRequest{
		Key: resources.NewKeyInt64(id, resources.TASKS),
		Attributes: resources.UpdateNftRequestAttributes{
			Status: status,
		},
	}

	endpoint := fmt.Sprintf("%s/%s/%s/%d", c.baseUrl, coreEndpoint, nftRequestEndpoint, id)
	requestAsBytes, err := json.Marshal(request)
	if err != nil {
		return errors.Wrap(err, "failed to marshal request")
	}

	return c.update(endpoint, requestAsBytes, nil)
}

func (c *Connector) ListNFTRequest(request models.ListNftRequestRequest) (*models.ListNftRequestResponse, error) {
	var result models.ListNftRequestResponse

	// setting full endpoint
	fullEndpoint := fmt.Sprintf("%s/%s/%s?%s", c.baseUrl, coreEndpoint, tasksEndpoint, urlval.MustEncode(request))

	// getting response
	if _, err := c.get(fullEndpoint, &result); err != nil {
		// errors are already wrapped
		return nil, err
	}

	return &result, nil
}
