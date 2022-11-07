package responses

import (
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/generator-svc/internal/data"
	"gitlab.com/tokend/nft-books/generator-svc/resources"
)

func NewGetTaskResponse(task data.Task, q data.BookQ) (*resources.TaskResponse, error) {
	var response resources.TaskResponse

	taskResource := task.Resource()
	taskResource.Relationships = getTaskRelationships(task)

	book, err := q.New().FilterByID(task.BookId).Get()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get a book")
	}

	bookResource, err := book.Resource()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get a book as a resource")
	}

	response.Data = taskResource
	response.Included.Add(bookResource)

	return &response, nil
}
