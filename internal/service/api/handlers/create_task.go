package handlers

import (
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/generator-svc/internal/data"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/helpers"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/requests"
	"gitlab.com/tokend/nft-books/generator-svc/resources"
	"net/http"
	"time"

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

	if !validateCreateTaskRequest(request, w, r) {
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

func validateCreateTaskRequest(request *requests.CreateTaskRequest, w http.ResponseWriter, r *http.Request) (ok bool) {
	database := helpers.GeneratorDB(r)
	restrictions := helpers.ApiRestrictions(r)
	statusFilter := resources.TaskFinishedGeneration

	tasks, err := database.Tasks().
		Sort(pgdb.Sorts{"created_at"}).
		Select(data.TaskSelector{
			Account: &request.Data.Attributes.Account,
			Status:  &statusFilter,
		})
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get select tasks from the database")
		ape.RenderErr(w, problems.InternalError())
		return false
	}

	tasksNumber := len(tasks)
	if uint64(tasksNumber) >= restrictions.MaxFailedAttempts {
		// TODO: make via jsonerrors.WithDetails
		ape.RenderErr(w, problems.BadRequest(errors.New("maximum attempts number exceeded"))[0])
		return false
	}

	var lastCreatedAt time.Time
	if tasksNumber > 0 {
		lastCreatedAt = tasks[tasksNumber-1].CreatedAt
	}

	durationAfterPreviousAttempt := time.Now().Sub(lastCreatedAt)
	if durationAfterPreviousAttempt < restrictions.RequestDelay {
		// TODO: make via jsonerrors.WithDetails
		ape.RenderErr(w, problems.BadRequest(errors.New("task delay not passed"))[0])
		return false

	}

	return true
}
