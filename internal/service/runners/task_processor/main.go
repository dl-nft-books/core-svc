package task_processor

import (
	"context"
	"strconv"

	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/running"
	s3connector "gitlab.com/tokend/nft-books/blob-svc/connector/api"
	"gitlab.com/tokend/nft-books/generator-svc/internal/config"
	"gitlab.com/tokend/nft-books/generator-svc/internal/data"
	"gitlab.com/tokend/nft-books/generator-svc/internal/data/postgres"
	"gitlab.com/tokend/nft-books/generator-svc/resources"
)

const cursorKey = "task_processor_cursor"

type TaskProcessor struct {
	name                string
	logger              *logan.Entry
	booksDB             data.BookQ
	generatorDB         data.GeneratorDB
	selector            data.TaskSelector
	runnerCfg           config.RunnerData
	signatureParams     *config.SignatureParams
	documenterConnector *s3connector.Connector
}

func New(cfg config.Config) *TaskProcessor {
	status := resources.TaskPending

	return &TaskProcessor{
		name: cfg.TaskProcessorCfg().Name,
		selector: data.TaskSelector{
			PageParams: pgdb.CursorPageParams{
				Cursor: cfg.TaskProcessorCfg().Cursor,
				Order:  pgdb.OrderTypeAsc,
				Limit:  cfg.TaskProcessorCfg().Limit,
			},
			Status: &status,
		},
		logger:              cfg.Log(),
		booksDB:             postgres.NewBooksQ(cfg.BookDB().DB),
		generatorDB:         postgres.NewGeneratorDB(cfg.GeneratorDB().DB),
		runnerCfg:           cfg.TaskProcessorCfg().Runner,
		signatureParams:     cfg.PdfSignatureParams(),
		documenterConnector: cfg.DocumenterConnector(),
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

func (p *TaskProcessor) run(ctx context.Context) error {
	return p.generatorDB.Transaction(func() error {
		tasks, err := p.getTasks(p.generatorDB)
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

			if err = p.generatorDB.Tasks().UpdateStatus(resources.TaskGenerating, task.Id); err != nil {
				return errors.Wrap(err, "failed to update task status", errFields)
			}
			if err = p.handleTask(task); err != nil {
				return errors.Wrap(err, "failed to handle task", errFields)
			}
			if err = p.generatorDB.Tasks().UpdateStatus(resources.TaskFinishedGeneration, task.Id); err != nil {
				return errors.Wrap(err, "failed to update task status", errFields)
			}
		}

		p.logger.Debugf("Successfully finished processing a batch of tasks (%d tasks)", len(tasks))
		return nil
	})
}

func (p *TaskProcessor) getTasks(db data.GeneratorDB) ([]data.Task, error) {
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
