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
	logger := helpers.Log(r)

	request, err := requests.NewListTasksRequest(r)
	if err != nil {
		logger.WithError(err).Error("invalid request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	tasks, err := helpers.GeneratorDB(r).Tasks().Select(data.TaskSelector{
		Account:      request.Account,
		Status:       request.Status,
		OffsetParams: &request.OffsetPageParams,
		IpfsHash:     request.IpfsHash,
	})
	if err != nil {
		logger.WithError(err).Error("unable to select tasks from database")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	taskListResponse, err := responses.NewTaskListResponse(r, request, tasks, helpers.BooksQ(r))
	if err != nil {
		logger.WithError(err).Error("unable to form task list response")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, *taskListResponse)
}
