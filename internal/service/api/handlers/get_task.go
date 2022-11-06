package handlers

import (
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/responses"
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/helpers"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/requests"
)

func GetTaskByID(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewTaskByIdRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	task, err := helpers.GeneratorDB(r).Tasks().GetById(request.Id)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get task")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if task == nil {
		ape.RenderErr(w, problems.NotFound())
		return
	}

	taskResponse, err := responses.NewGetTaskResponse(*task, helpers.BooksQ(r))
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get task response")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, *taskResponse)
}
