package connector

import (
	"fmt"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"github.com/dl-nft-books/core-svc/connector/models"
)

const (
	promocodesEndpoint = "promocodes"
)

func (c *Connector) GetPromocodeById(id int64) (*models.PromocodeResponse, error) {
	var promocode models.PromocodeResponse

	// setting full endpoint
	fullEndpoint := fmt.Sprintf("%s/%s/%d", c.baseUrl, generatorEndpoint, id)

	// getting response
	_, err := c.get(fullEndpoint, &promocode)
	if err != nil {
		// errors are already wrapped
		return nil, errors.From(err, logan.F{"id": id})
	}
	return &promocode, nil
}

func (c *Connector) RollbackPromocode(id int64) error {
	// setting full endpoint
	fullEndpoint := fmt.Sprintf("%s/%s/%s/%s/%d", c.baseUrl, generatorEndpoint, promocodesEndpoint, "rollback", id)

	// getting response
	err := c.update(fullEndpoint, []byte{}, nil)
	if err != nil {
		// errors are already wrapped
		return errors.From(err, logan.F{"id": id})
	}
	return nil
}
