package handlers

import (
	"github.com/dl-nft-books/core-svc/internal/data"
	"github.com/dl-nft-books/core-svc/internal/data/postgres"
	"github.com/dl-nft-books/core-svc/internal/service/api/helpers"
	"github.com/dl-nft-books/core-svc/internal/service/api/jsonerrors"
	"github.com/dl-nft-books/core-svc/internal/service/api/requests"
	"github.com/dl-nft-books/core-svc/resources"
	"net/http"
	"time"

	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3/errors"

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
	getBookResponse, err := helpers.Booker(r).GetBookById(bookId, request.Data.Attributes.ChainId)
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

	// Validating info
	banner := request.Data.Attributes.Banner

	if err = helpers.CheckBannerMimeType(banner.Attributes.MimeType, r); err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	// Setting banner link
	if err = helpers.SetMediaLink(r, &banner); err != nil {
		helpers.Log(r).WithError(err).Error("failed to set banner link")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	media := helpers.MarshalMedia(&banner)
	if media == nil {
		helpers.Log(r).Error("failed to marshal media")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	// Then creating task
	createdTaskId, err := helpers.DB(r).Tasks().Insert(data.Task{
		BookId:    bookId,
		Banner:    media[0],
		Account:   request.Data.Attributes.Account,
		Status:    resources.TaskPending,
		CreatedAt: time.Now(),
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

	// Validating if a user has generated too many unpaid books
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

	// If no tasks were found - simply return that everything is ok
	if tasksNumber == 0 {
		return true
	}

	if uint64(tasksNumber) >= restrictions.MaxFailedAttempts {
		ape.RenderErr(w, jsonerrors.WithDetails(
			problems.BadRequest(errors.New("max amount of tries exceeded"))[0],
			jsonerrors.ApiMaxTriesExceeded,
		))
		return false
	}

	// Validating how often user tries to buy a book
	lastCreatedAt := tasks[tasksNumber-1].CreatedAt

	durationAfterPreviousAttempt := time.Now().Sub(lastCreatedAt)
	if durationAfterPreviousAttempt < restrictions.RequestDelay {
		ape.RenderErr(w, jsonerrors.WithDetails(
			problems.BadRequest(errors.New("task delay not passed"))[0],
			jsonerrors.TaskDelayNotPassed,
		))
		return false
	}

	return true
}
