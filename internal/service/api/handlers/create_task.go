package handlers

import (
	"net/http"

	"gitlab.com/tokend/nft-books/generator-svc/internal/data"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/helpers"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/requests"
	"gitlab.com/tokend/nft-books/generator-svc/resources"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

type KeyResponse struct {
	Data resources.Key `json:"data"`
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	logger := helpers.Log(r)

	request, err := requests.NewCreateTaskRequest(r)
	if err != nil {
		logger.WithError(err).Error("failed to fetch create task request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	bookId := int64(request.Data.Attributes.BookId)

	// Check if book exists
	book, err := helpers.BooksQ(r).FilterActual().FilterByID(bookId).Get()
	if err != nil {
		logger.WithError(err).Errorf("failed to get book with id #%v", bookId)
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if book == nil {
		ape.RenderErr(w, problems.NotFound())
		return
	}

	// Then creating task
	createdTaskId, err := helpers.GeneratorDB(r).Tasks().Insert(data.Task{
		BookId:    bookId,
		Signature: request.Data.Attributes.Signature,
		Account:   request.Data.Attributes.Account,
		Status:    resources.TaskPending,
	})
	if err != nil {
		logger.WithError(err).Errorf("failed to create new task with book id #%v", bookId)
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, KeyResponse{Data: resources.NewKeyInt64(createdTaskId, resources.TASKS)})
}
