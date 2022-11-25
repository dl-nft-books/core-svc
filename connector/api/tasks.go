package api

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/generator-svc/resources"
)

const tasksEndpoint = "tasks"

func (c *Connector) CreateTask(request resources.CreateTask) (id int64, err error) {
	var response resources.KeyResponse

	endpoint := fmt.Sprintf("%s/%s", c.baseUrl, tasksEndpoint)
	requestAsBytes, err := json.Marshal(request)
	if err != nil {
		return 0, errors.Wrap(err, "failed to marshal request")
	}

	if err = c.post(endpoint, requestAsBytes, &response); err != nil {
		return 0, errors.Wrap(err, "failed to create token")
	}

	createdTokenId := cast.ToInt64(response.Data.ID)
	return createdTokenId, nil
}

func (c *Connector) UpdateTask(request resources.UpdateTask) error {
	endpoint := fmt.Sprintf("%s/%s/%s", c.baseUrl, tasksEndpoint, request.ID)
	requestAsBytes, err := json.Marshal(request)
	if err != nil {
		return errors.Wrap(err, "failed to marshal request")
	}

	return c.update(endpoint, requestAsBytes, nil)
}
