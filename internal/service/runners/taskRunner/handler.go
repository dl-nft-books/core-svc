package taskRunner

import "gitlab.com/tokend/nft-books/generator-svc/internal/data"

func (r *TaskRunner) handleTask(task data.Task) error {
	r.logger.Debugf("Started processing task with id of %d", task.Id)
	// TODO: Add implementation here

	r.logger.Debugf("Successfully finished processing task with id of %d", task.Id)
	return nil
}
