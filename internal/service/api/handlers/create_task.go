package handlers

import (
	"gitlab.com/tokend/nft-books/generator-svc/internal/data"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/helpers"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/requests"
	"gitlab.com/tokend/nft-books/generator-svc/resources"
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewCreateTaskRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	createdTaskId, err := helpers.GeneratorDB(r).Tasks().Insert(data.Task{
		BookId:    int64(request.Data.Attributes.BookId),
		Signature: request.Data.Attributes.Signature,
		Account:   request.Data.Attributes.Account,
		Status:    resources.TaskPending,
	})

	ape.Render(w, resources.NewKeyInt64(createdTaskId, resources.TASKS))
}
