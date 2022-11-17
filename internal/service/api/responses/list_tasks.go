package responses

import (
	"gitlab.com/tokend/nft-books/generator-svc/internal/data/external"
	"net/http"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/generator-svc/internal/data"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/requests"
	"gitlab.com/tokend/nft-books/generator-svc/resources"
)

func NewTaskListResponse(r *http.Request, request *requests.ListTasksRequest, tasks []data.Task, q external.BookQ) (*resources.TaskListResponse, error) {
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

		book, err := q.New().FilterByID(task.BookId).Get()
		if err != nil {
			return nil, errors.Wrap(err, "failed to get book by its id", logan.F{
				"book_id": book.ID,
			})
		}

		bookResource, err := book.Resource()
		if err != nil {
			return nil, errors.Wrap(err, "failed to convert book to the resource format")
		}

		response.Included.Add(bookResource)
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
