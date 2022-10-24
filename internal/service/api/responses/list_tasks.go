package responses

import (
	"gitlab.com/tokend/nft-books/generator-svc/internal/data"
	"gitlab.com/tokend/nft-books/generator-svc/resources"
)

func NewTaskListResponse(tasks []data.Task) (response resources.TaskListResponse) {
	if len(tasks) == 0 {
		return resources.TaskListResponse{
			Data: []resources.Task{},
		}
	}

	for _, task := range tasks {
		response.Data = append(response.Data, task.Resource())
	}

	return response
}
