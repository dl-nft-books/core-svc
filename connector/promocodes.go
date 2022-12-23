package connector

import (
	"fmt"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/generator-svc/connector/models"
)

const (
	promocodesEndpoint = "promocodes"
)

func (c *Connector) GetPromocodeById(id int64) (*models.PromocodeResponse, error) {
	var promocode models.PromocodeResponse

	// setting full endpoint
	fullEndpoint := fmt.Sprintf("%s/%s/%d", c.baseUrl, generatorEndpoint, id)

	// getting response
	found, err := c.get(fullEndpoint, &promocode)
	if err != nil {
		// errors are already wrapped
		return nil, errors.From(err, logan.F{"id": id})
	}
	if !found {
		return nil, nil
	}

	return &promocode, nil
}

func (c *Connector) RollbackPromocode(id int64) error {
	var res int64

	// setting full endpoint
	fullEndpoint := fmt.Sprintf("%s/%s/%s/%s/%d", c.baseUrl, generatorEndpoint, promocodesEndpoint, "rollback", id)

	// getting response
	found, err := c.get(fullEndpoint, &res)
	if err != nil {
		// errors are already wrapped
		return errors.From(err, logan.F{"id": id})
	}
	if !found {
		return nil
	}

	return nil
}
