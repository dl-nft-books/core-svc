package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokend/nft-books/generator-svc/internal/data"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/helpers"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/requests"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/responses"
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

	ape.Render(w, responses.NewTaskListResponse(tasks))
}
