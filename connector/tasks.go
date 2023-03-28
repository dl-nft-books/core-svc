package connector

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/urlval"
	"github.com/dl-nft-books/core-svc/connector/models"
	"github.com/dl-nft-books/core-svc/resources"
)

const (
	generatorEndpoint = "generator"
	tasksEndpoint     = "tasks"
)

func (c *Connector) CreateTask(params models.CreateTaskParams) (id int64, err error) {
	var (
		response resources.KeyResponse
		request  = resources.CreateTaskRequest{
			Data: resources.CreateTask{
				Key: resources.NewKeyInt64(0, resources.TASKS),
				Attributes: resources.CreateTaskAttributes{
					Account:   params.Account,
					BookId:    params.BookId,
					Signature: params.Signature,
				},
			},
			Included: resources.Included{},
		}
	)

	endpoint := fmt.Sprintf("%s/%s/%s", c.baseUrl, generatorEndpoint, tasksEndpoint)
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

func (c *Connector) UpdateTask(params models.UpdateTaskParams) error {
	request := resources.UpdateTaskRequest{
		Data: resources.UpdateTask{
			Key: resources.NewKeyInt64(params.Id, resources.TASKS),
			Attributes: resources.UpdateTaskAttributes{
				Status:  params.Status,
				TokenId: params.TokenId,
			},
		},
		Included: resources.Included{},
	}

	endpoint := fmt.Sprintf("%s/%s/%s/%s", c.baseUrl, generatorEndpoint, tasksEndpoint, request.Data.Key.ID)
	requestAsBytes, err := json.Marshal(request)
	if err != nil {
		return errors.Wrap(err, "failed to marshal request")
	}

	return c.update(endpoint, requestAsBytes, nil)
}

func (c *Connector) ListTasks(request models.ListTasksRequest) (*models.ListTasksResponse, error) {
	var result models.ListTasksResponse

	// setting full endpoint
	fullEndpoint := fmt.Sprintf("%s/%s/%s?%s", c.baseUrl, generatorEndpoint, tasksEndpoint, urlval.MustEncode(request))

	// getting response
	if _, err := c.get(fullEndpoint, &result); err != nil {
		// errors are already wrapped
		return nil, err
	}

	return &result, nil
}

func (c *Connector) GetTaskById(id int64) (*models.TaskResponse, error) {
	var result models.TaskResponse

	// setting full endpoint
	fullEndpoint := fmt.Sprintf("%s/%s/%s/%d", c.baseUrl, generatorEndpoint, tasksEndpoint, id)

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
