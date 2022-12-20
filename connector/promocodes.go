package connector

import (
	"encoding/json"
	"fmt"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/generator-svc/connector/models"
	"gitlab.com/tokend/nft-books/generator-svc/resources"
)

const (
	promocodesRollbackEndpoint = "promocodes/rollback"
)

func (c *Connector) GetPromocodeById(id int64) (*models.PromocodeResponse, error) {
	var promocode models.PromocodeResponse

	// setting full endpoint
	fullEndpoint := fmt.Sprintf("%s/%s/%s/%d", c.baseUrl, generatorEndpoint, promocodesRollbackEndpoint, id)

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
	var promocode, err = c.GetPromocodeById(id)

	if err != nil {
		// errors are already wrapped
		return errors.From(err, logan.F{"id": id})
	}
	if promocode == nil {
		return nil
	}

	*promocode.Data.Attributes.LeftUsages--
	request := resources.UpdatePromocodeRequest{
		Data: resources.UpdatePromocode{
			Key: resources.NewKeyInt64(id, resources.PROMOCODE),
			Attributes: resources.UpdatePromocodeAttributes{
				LeftUsages: promocode.Data.Attributes.LeftUsages,
			},
		},
		Included: resources.Included{},
	}
	endpoint := fmt.Sprintf("%s/%s/%s/%s", c.baseUrl, generatorEndpoint, promocodesRollbackEndpoint, id)
	requestAsBytes, err := json.Marshal(request)
	if err != nil {
		return errors.Wrap(err, "failed to marshal request")
	}

	return c.update(endpoint, requestAsBytes, nil)
}
