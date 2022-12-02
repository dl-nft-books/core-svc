package handlers

import (
	"net/http"
	"time"

	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/generator-svc/internal/data"
	"gitlab.com/tokend/nft-books/generator-svc/internal/data/postgres"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/helpers"
	"gitlab.com/tokend/nft-books/generator-svc/internal/service/api/requests"
	"gitlab.com/tokend/nft-books/generator-svc/resources"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewCreateTaskRequest(r)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to fetch create task request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	if !validateCreateTaskRequest(request, w, r) {
		return
	}

	bookId := request.Data.Attributes.BookId

	// Check if book exists
	getBookResponse, err := helpers.Booker(r).GetBookById(bookId)
	if err != nil {
		helpers.Log(r).WithError(err).Errorf("failed to get book with id #%v", bookId)
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if getBookResponse == nil {
		helpers.Log(r).Info("corresponding book was not found")
		ape.RenderErr(w, problems.NotFound())
		return
	}

	// Then creating task
	createdTaskId, err := helpers.DB(r).Tasks().Insert(data.Task{
		BookId:    bookId,
		Signature: request.Data.Attributes.Signature,
		Account:   request.Data.Attributes.Account,
		Status:    resources.TaskPending,
	})
	if err != nil {
		helpers.Log(r).WithError(err).Errorf("failed to create new task with book id #%v", bookId)
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, resources.KeyResponse{
		Data: resources.NewKeyInt64(createdTaskId, resources.TASKS),
	})
}

func validateCreateTaskRequest(request *requests.CreateTaskRequest, w http.ResponseWriter, r *http.Request) (ok bool) {
	var (
		database     = helpers.DB(r)
		restrictions = helpers.ApiRestrictions(r)
		statusFilter = resources.TaskFinishedGeneration
	)

	// validating if user have generated a lot of books and did not pay for them
	tasks, err := database.Tasks().
		Sort(pgdb.Sorts{postgres.TasksCreatedAt}).
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

	//  if no tasks found -- simply return that everything is ok
	if tasksNumber == 0 {
		return true
	}

	if uint64(tasksNumber) >= restrictions.MaxFailedAttempts {
		// TODO: make via jsonerrors.WithDetails
		ape.RenderErr(w, problems.BadRequest(errors.New("maximum attempts number exceeded"))[0])
		return false
	}

	// then validating how often user try to buy the book
	lastCreatedAt := tasks[tasksNumber-1].CreatedAt

	durationAfterPreviousAttempt := time.Now().Sub(lastCreatedAt)
	if durationAfterPreviousAttempt < restrictions.RequestDelay {
		// TODO: make via jsonerrors.WithDetails
		ape.RenderErr(w, problems.BadRequest(errors.New("task delay not passed"))[0])
		return false
	}

	return true
}
