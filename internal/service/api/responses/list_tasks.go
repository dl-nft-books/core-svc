package responses

import (
	booker "gitlab.com/tokend/nft-books/book-svc/connector"
	"net/http"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/generator-svc/internal/data"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/requests"
	"gitlab.com/tokend/nft-books/generator-svc/resources"
)

func NewTaskListResponse(r *http.Request, request *requests.ListTasksRequest, tasks []data.Task, booksApi booker.Connector) (*resources.TaskListResponse, error) {
	response := resources.TaskListResponse{}

	if len(tasks) == 0 {
		return &resources.TaskListResponse{
			Data: make([]resources.Task, 0),
		}, nil
	}

	for _, task := range tasks {
		taskResource := task.Resource()
		taskResource.Relationships = getTaskRelationships(task)

		response.Data = append(response.Data, taskResource)

		bookResponse, err := booksApi.GetBookById(task.BookId)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get book by its id", logan.F{
				"book_id": task.BookId,
			})
		}

		response.Included.Add(convertBookToResource(*bookResponse))
	}

	response.Links = requests.GetOffsetLinksWithSort(r, request.OffsetPageParams, request.Sorts)
	return &response, nil
}

func getTaskRelationships(task data.Task) resources.TaskRelationships {
	bookKey := resources.NewKeyInt64(task.BookId, resources.BOOKS)

	return resources.TaskRelationships{
		Book: &resources.Relation{
			Data:  &bookKey,
			Links: nil,
		},
	}
}
