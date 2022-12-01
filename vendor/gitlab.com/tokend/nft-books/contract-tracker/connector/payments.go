package connector

import (
	"fmt"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/urlval"
	models "gitlab.com/tokend/nft-books/contract-tracker/connector/models"
)

const trackerEndpoint = "tracker"

func (c *Connector) ListPayments(request models.ListPaymentsParams) (*models.ListPaymentsResponse, error) {
	var result models.ListPaymentsResponse

	// setting full endpoint
	fullEndpoint := fmt.Sprintf("%s/%s?%s", c.baseUrl, trackerEndpoint, urlval.MustEncode(request))

	// getting response
	if _, err := c.get(fullEndpoint, &result); err != nil {
		// errors are already wrapped
		return nil, err
	}

	return &result, nil
}

func (c *Connector) GetPaymentById(id int64) (*models.GetPaymentResponse, error) {
	var result models.GetPaymentResponse

	// setting full endpoint
	fullEndpoint := fmt.Sprintf("%s/%s/%d", c.baseUrl, trackerEndpoint, id)

	// getting response
	found, err := c.get(fullEndpoint, &result)
	if err != nil {
		// errors are already wrapped
		return nil, errors.From(err, logan.F{"id": id})
	}
	if !found {
		return nil, nil
	}

	return &result, nil
}
