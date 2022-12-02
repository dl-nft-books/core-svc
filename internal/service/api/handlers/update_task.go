package handlers

import (
	"net/http"

	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokend/nft-books/generator-svc/internal/data"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/helpers"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/requests"
	"gitlab.com/tokend/nft-books/generator-svc/resources"
)

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	// Getting the update token request
	request, err := requests.NewUpdateTaskRequest(r)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to fetch update task request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	tasksQ := helpers.DB(r).Tasks()

	// Validating whether specified task exists
	taskId := cast.ToInt64(request.Data.ID)
	task, err := tasksQ.GetById(taskId)
	if err != nil {
		helpers.Log(r).WithError(err).Errorf("failed to get a task with id of %v", taskId)
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if task == nil {
		ape.RenderErr(w, problems.NotFound())
		return
	}

	// Clearing selector filters and applying updator ones
	tasksQ = applyTaskUpdateFilters(tasksQ.New(), *request)

	if err = tasksQ.Update(taskId); err != nil {
		helpers.Log(r).WithError(err).Errorf("failed to update task with id of #%v", taskId)
		ape.RenderErr(w, problems.InternalError())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func applyTaskUpdateFilters(q data.TasksQ, request resources.UpdateTaskRequest) data.TasksQ {
	if request.Data.Attributes.TokenId != nil {
		q = q.UpdateTokenId(*request.Data.Attributes.TokenId)
	}
	if request.Data.Attributes.Status != nil {
		q = q.UpdateStatus(*request.Data.Attributes.Status)
	}

	return q
}
