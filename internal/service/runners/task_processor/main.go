package task_processor

import (
	"context"
	"fmt"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/running"
	documenter "gitlab.com/tokend/nft-books/blob-svc/connector/api"
	booker "gitlab.com/tokend/nft-books/book-svc/connector"
	"gitlab.com/tokend/nft-books/generator-svc/internal/config"
	"gitlab.com/tokend/nft-books/generator-svc/internal/data"
	"gitlab.com/tokend/nft-books/generator-svc/internal/data/postgres"
	"gitlab.com/tokend/nft-books/generator-svc/resources"
	"net/http"
	"strconv"
	"time"
)

const cursorKey = "task_processor_cursor"

type TaskProcessor struct {
	name     string
	logger   *logan.Entry
	db       data.DB
	selector data.TaskSelector

	runnerCfg       config.RunnerData
	cleanerCfg      config.CleanerData
	signatureParams *config.SignatureParams

	booksApi   *booker.Connector
	documenter *documenter.Connector
}

func New(cfg config.Config) *TaskProcessor {
	status := resources.TaskPending

	return &TaskProcessor{
		name:   cfg.TaskProcessorCfg().Name,
		db:     postgres.NewDB(cfg.DB()),
		logger: cfg.Log(),

		selector: data.TaskSelector{
			PageParams: &pgdb.CursorPageParams{
				Cursor: cfg.TaskProcessorCfg().Cursor,
				Order:  pgdb.OrderTypeAsc,
				Limit:  cfg.TaskProcessorCfg().Limit,
			},
			Status: &status,
		},

		runnerCfg:       cfg.TaskProcessorCfg().Runner,
		cleanerCfg:      cfg.TaskProcessorCfg().Cleaner,
		signatureParams: cfg.PdfSignatureParams(),

		booksApi:   cfg.BookerConnector(),
		documenter: cfg.DocumenterConnector(),
	}
}

func (p *TaskProcessor) Run(ctx context.Context) {
	running.WithBackOff(
		ctx, p.logger,
		p.name, p.run,
		p.runnerCfg.NormalPeriod,
		p.runnerCfg.MinAbnormalPeriod,
		p.runnerCfg.MaxAbnormalPeriod,
	)
}

func (p *TaskProcessor) cleanTasks() error {
	unresTasks, err := p.getUnresolvedTasks()
	if err != nil {
		return errors.Wrap(err, "failed to get tasks from the database")
	}

	if len(unresTasks) == 0 {
		p.logger.Debug("Found no unresolved tasks to process")
		return nil
	}

	for _, task := range unresTasks {
		errFields := logan.F{
			"task_id":        task.Id,
			"task_signature": task.Signature,
			"task_status":    task.Status,
		}

		fileName := fmt.Sprintf("%s.pdf", task.FileIpfsHash)

		statusCode, err := p.documenter.DeleteDocument(fileName)

		if err != nil {
			return errors.Wrap(err, "failed to delete document from S3", errFields)
		}

		if statusCode != http.StatusOK {
			return errors.New("failed to delete document from S3")
		}

		if err := p.db.New().Tasks().Delete(task.Id); err != nil {
			return errors.Wrap(err, "failed to delete task from data base", errFields)
		}

		p.logger.Debug("Document deleted from S3")

	}

	return nil
}

func (p *TaskProcessor) run(ctx context.Context) error {
	return p.db.Transaction(func() error {
		if err := p.cleanTasks(); err != nil {
			return errors.Wrap(err, "failed to clean tasks")
		}

		tasks, err := p.getTasks(p.db)
		if err != nil {
			return errors.Wrap(err, "failed to get tasks from the database")
		}
		if len(tasks) == 0 {
			p.logger.Debug("Found no tasks to process")
			return nil
		}
		p.logger.Debugf("Found %d task(s) to process", len(tasks))

		for _, task := range tasks {
			errFields := logan.F{
				"task_id":        task.Id,
				"task_signature": task.Signature,
				"task_status":    task.Status,
			}

			if err = p.db.Tasks().UpdateStatus(resources.TaskGenerating).Update(task.Id); err != nil {
				return errors.Wrap(err, "failed to update task status", errFields)
			}

			if err = p.handleTask(task); err != nil {
				return errors.Wrap(err, "failed to handle task", errFields)
			}

			if err = p.db.Tasks().UpdateStatus(resources.TaskFinishedGeneration).Update(task.Id); err != nil {
				return errors.Wrap(err, "failed to update task status", errFields)
			}
		}

		p.logger.Debugf("Successfully finished processing a batch of tasks (%d tasks)", len(tasks))
		return nil
	})
}

func (p *TaskProcessor) getUnresolvedTasks() ([]data.Task, error) {
	waitingPeriod := time.Duration(p.cleanerCfg.MaxWaitingPeriod)
	status := resources.TaskFinishedGeneration
	selector := data.TaskSelector{
		Status: &status,
	}

	tasks, err := p.db.New().Tasks().FilterByMaxWaitingPeriod(waitingPeriod).Select(selector)

	if err != nil {
		return nil, errors.Wrap(err, "failed to get subtasks from db")
	}

	return tasks, nil
}

func (p *TaskProcessor) getTasks(db data.DB) ([]data.Task, error) {
	cursorKV, err := db.KeyValue().LockingGet(cursorKey)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get current cursor value")
	}

	if cursorKV == nil {
		cursorKV = &data.KeyValue{
			Key:   cursorKey,
			Value: "0",
		}
	}

	cursor, err := strconv.ParseInt(cursorKV.Value, 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse cursor")
	}

	p.selector.PageParams.Cursor = uint64(cursor)
	tasks, err := db.Tasks().Select(p.selector)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get subtasks from db")
	}
	if len(tasks) == 0 {
		return nil, nil
	}

	return tasks, db.KeyValue().Upsert(
		data.KeyValue{
			Key:   cursorKey,
			Value: strconv.FormatInt(tasks[len(tasks)-1].Id, 10),
		},
	)
}
