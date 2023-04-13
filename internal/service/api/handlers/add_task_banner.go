package handlers

import (
	"bytes"
	"github.com/dl-nft-books/core-svc/internal/service/api/responses"
	"github.com/dl-nft-books/core-svc/resources"
	"io"
	"net/http"

	"github.com/dl-nft-books/core-svc/internal/service/api/helpers"
	"github.com/dl-nft-books/core-svc/internal/service/api/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func AddTaskBanner(w http.ResponseWriter, r *http.Request) {
	logger := helpers.Log(r)

	id, file, err := requests.NewAddTaskBannerRequest(r)
	if err != nil {
		logger.WithError(err).Error("failed to fetch get task request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	task, err := helpers.DB(r).Tasks().GetById(id)
	if err != nil {
		logger.WithError(err).Error("failed to get task")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if task == nil {
		logger.WithError(err).Error("task with such id doesn't exist")
		ape.RenderErr(w, problems.NotFound())
		return
	}
	// Read file contents into buffer
	var buf bytes.Buffer
	_, err = io.Copy(&buf, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Convert buffer to []byte slice
	fileBytes := buf.Bytes()

	if err = helpers.DB(r).Tasks().UpdateStatus(resources.TaskGenerating).Update(task.Id); err != nil {
		logger.WithError(err).Error("failed to update task status")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	err = helpers.HandleTask(r, logger, *task, fileBytes)
	if err != nil {
		logger.WithError(err).Error("failed to handle task")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err = helpers.DB(r).Tasks().UpdateStatus(resources.TaskFinishedGeneration).Update(task.Id); err != nil {
		logger.WithError(err).Error("failed to update task status")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	task, err = helpers.DB(r).Tasks().GetById(id)
	if err != nil {
		logger.WithError(err).Error("failed to get task")
		ape.RenderErr(w, problems.InternalError())
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
