package responses

import (
	booker "github.com/dl-nft-books/book-svc/connector"
	"github.com/dl-nft-books/core-svc/internal/data"
	"github.com/dl-nft-books/core-svc/resources"
	"gitlab.com/distributed_lab/logan/v3/errors"
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
	response.Included.Add(convertBookToResource(*bookResponse))

	return &response, nil
}
