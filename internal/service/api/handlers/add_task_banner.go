package handlers

import (
	"bytes"
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
	if len(task.Banner) > 0 {
		logger.WithError(err).Error("banner is already uploaded for this task")
		ape.RenderErr(w, problems.Conflict())
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
	if err = helpers.DB(r).Tasks().UpdateBanner(fileBytes).Update(id); err != nil {
		logger.WithError(err).Error("failed to add task banner")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
