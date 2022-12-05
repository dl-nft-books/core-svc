package connector

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/urlval"
	"gitlab.com/tokend/nft-books/generator-svc/connector/models"
	"gitlab.com/tokend/nft-books/generator-svc/resources"
)

const tasksEndpoint = "tasks"

func (c *Connector) CreateTask(params models.CreateTaskParams) (id int64, err error) {
	var (
		response resources.KeyResponse
		request  = resources.CreateTask{
			Key: resources.NewKeyInt64(0, resources.TASKS),
			Attributes: resources.CreateTaskAttributes{
				Account:   params.Account,
				BookId:    params.BookId,
				Signature: params.Signature,
			},
		}
	)

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

func (c *Connector) UpdateTask(params models.UpdateTaskParams) error {
	request := resources.UpdateTask{
		Key: resources.NewKeyInt64(params.Id, resources.TASKS),
		Attributes: resources.UpdateTaskAttributes{
			Status:  params.Status,
			TokenId: params.TokenId,
		},
	}

	endpoint := fmt.Sprintf("%s/%s/%s", c.baseUrl, tasksEndpoint, request.ID)
	requestAsBytes, err := json.Marshal(request)
	if err != nil {
		return errors.Wrap(err, "failed to marshal request")
	}

	return c.update(endpoint, requestAsBytes, nil)
}

func (c *Connector) ListTasks(request models.ListTasksRequest) (*models.ListTasksResponse, error) {
	var result models.ListTasksResponse

	// setting full endpoint
	fullEndpoint := fmt.Sprintf("%s/%s?%s", c.baseUrl, tasksEndpoint, urlval.MustEncode(request))

	// getting response
	if err := c.get(fullEndpoint, &result); err != nil {
		// errors are already wrapped
		return nil, err
	}

	return &result, nil
}

func (c *Connector) GetTaskById(id int64) (*models.TaskResponse, error) {
	var result models.TaskResponse

	// setting full endpoint
	fullEndpoint := fmt.Sprintf("%s/%s/%d", c.baseUrl, tasksEndpoint, id)

	// getting response
	if err := c.get(fullEndpoint, &result); err != nil {
		// errors are already wrapped
		return nil, errors.From(err, logan.F{"id": id})
	}

	return &result, nil
}
