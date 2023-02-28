package handlers

import (
	"net/http"

	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/responses"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/helpers"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/requests"
)

func GetTaskByID(w http.ResponseWriter, r *http.Request) {
	logger := helpers.Log(r)

	request, err := requests.NewTaskByIdRequest(r)
	if err != nil {
		logger.WithError(err).Error("failed to fetch get task request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	task, err := helpers.DB(r).Tasks().GetById(request.Id)
	if err != nil {
		logger.WithError(err).Error("failed to get task")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if task == nil {
		ape.RenderErr(w, problems.NotFound())
		return
	}

	taskResponse, err := responses.NewGetTaskResponse(*task, helpers.Booker(r))
	if err != nil {
		logger.WithError(err).Error("failed to get task response")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, *taskResponse)
}
