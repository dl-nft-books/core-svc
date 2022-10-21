package task_processor

import "gitlab.com/tokend/nft-books/generator-svc/internal/data"

func (p *TaskProcessor) handleTask(task data.Task) error {
	p.logger.Debugf("Started processing task with id of %d", task.Id)
	// TODO: Add implementation here

	p.logger.Debugf("Successfully finished processing task with id of %d", task.Id)
	return nil
}
