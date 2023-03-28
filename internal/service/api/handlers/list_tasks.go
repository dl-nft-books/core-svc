package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"github.com/dl-nft-books/core-svc/internal/data"
	"github.com/dl-nft-books/core-svc/internal/service/api/helpers"
	"github.com/dl-nft-books/core-svc/internal/service/api/requests"
	"github.com/dl-nft-books/core-svc/internal/service/api/responses"
)

func ListTasks(w http.ResponseWriter, r *http.Request) {
	logger := helpers.Log(r)

	request, err := requests.NewListTasksRequest(r)
	if err != nil {
		logger.WithError(err).Error("failed to fetch list tasks request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	tasks, err := helpers.DB(r).Tasks().Select(data.TaskSelector{
		Account:      request.Account,
		Status:       request.Status,
		OffsetParams: &request.OffsetPageParams,
		IpfsHash:     request.IpfsHash,
		TokenId:      request.TokenId,
	})
	if err != nil {
		logger.WithError(err).Error("unable to select tasks from database")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	taskListResponse, err := responses.NewTaskListResponse(r, request, tasks, *helpers.Booker(r))
	if err != nil {
		logger.WithError(err).Error("unable to form task list response")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, *taskListResponse)
}
