package taskRunner

import (
	"context"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/running"
	"gitlab.com/tokend/nft-books/generator-svc/internal/config"
	"gitlab.com/tokend/nft-books/generator-svc/internal/data"
	"gitlab.com/tokend/nft-books/generator-svc/internal/data/postgres"
	"gitlab.com/tokend/nft-books/generator-svc/resources"
	"strconv"
)

const cursorKey = "task_runner_cursor"

type TaskRunner struct {
	name   string
	logger *logan.Entry

	db       data.DB
	selector data.TaskSelector

	runnerCfg config.RunnerData
}

func New(cfg config.Config) *TaskRunner {
	status := resources.TaskPending

	return &TaskRunner{
		name: cfg.TaskRunnerCfg().Name,
		selector: data.TaskSelector{
			PageParams: pgdb.CursorPageParams{
				Cursor: cfg.TaskRunnerCfg().Cursor,
				Order:  pgdb.OrderTypeAsc,
				Limit:  cfg.TaskRunnerCfg().Limit,
			},
			Status: &status,
		},
		logger:    cfg.Log(),
		db:        postgres.NewDB(cfg.DB()),
		runnerCfg: cfg.TaskRunnerCfg().Runner,
	}
}

func (r *TaskRunner) Run(ctx context.Context) {
	running.WithBackOff(
		ctx, r.logger,
		r.name, r.run,
		r.runnerCfg.NormalPeriod,
		r.runnerCfg.MinAbnormalPeriod,
		r.runnerCfg.MaxAbnormalPeriod,
	)
}

func (r *TaskRunner) run(ctx context.Context) error {
	return r.db.Transaction(func() error {
		tasks, err := r.getTasks(r.db)
		if err != nil {
			return errors.Wrap(err, "failed to get tasks from the database")
		}
		if len(tasks) == 0 {
			r.logger.Debug("Found no tasks to process")
			return nil
		}
		r.logger.Debugf("Found %d task(s) to process", len(tasks))

		for _, task := range tasks {
			errFields := logan.F{
				"task_id":        task.Id,
				"task_signature": task.Signature,
				"task_status":    task.Status,
			}

			if err = r.db.Tasks().UpdateStatus(resources.TaskGenerating, task.Id); err != nil {
				return errors.Wrap(err, "failed to update task status", errFields)
			}
			if err = r.handleTask(task); err != nil {
				return errors.Wrap(err, "failed to handle task", errFields)
			}
			if err = r.db.Tasks().UpdateStatus(resources.TaskFinishedGeneration, task.Id); err != nil {
				return errors.Wrap(err, "failed to update task status", errFields)
			}
		}

		r.logger.Debugf("Successfully finished processing a batch of tasks", len(tasks))
		return nil
	})
}

func (r *TaskRunner) getTasks(db data.DB) ([]data.Task, error) {
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

	r.selector.PageParams.Cursor = uint64(cursor)
	tasks, err := db.Tasks().Select(r.selector)
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
