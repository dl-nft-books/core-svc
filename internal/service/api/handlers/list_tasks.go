package handlers

import (
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokend/nft-books/generator-svc/internal/data"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/helpers"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/requests"
	"gitlab.com/tokend/nft-books/generator-svc/resources"
	"net/http"
)

func ListTasks(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewListTasksRequest(r)
	if err != nil {
		helpers.Log(r).WithError(err).Error("invalid request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	tasks, err := helpers.GeneratorDB(r).Tasks().Select(data.TaskSelector{
		Account: request.Account,
		Status:  request.Status,
	})
	if err != nil {
		helpers.Log(r).WithError(err).Error("unable to select tasks from database")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, formTaskListResponse(tasks))
}

func formTaskListResponse(tasks []data.Task) (response resources.TaskListResponse) {
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
