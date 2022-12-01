package responses

import (
	"gitlab.com/distributed_lab/logan/v3/errors"
	booker "gitlab.com/tokend/nft-books/book-svc/connector"
	"gitlab.com/tokend/nft-books/generator-svc/internal/data"
	"gitlab.com/tokend/nft-books/generator-svc/resources"
)

func NewGetTaskResponse(task data.Task, booker *booker.Connector) (*resources.TaskResponse, error) {
	var response resources.TaskResponse

	taskResource := task.Resource()
	taskResource.Relationships = getTaskRelationships(task)

	bookResponse, err := booker.GetBookById(task.BookId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get a book")
	}

	response.Data = taskResource
	bookResponseData := bookResponse.Data
	response.Included.Add(&bookResponseData)

	return &response, nil
}
