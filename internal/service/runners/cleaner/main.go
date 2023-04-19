package cleaner

import (
	"context"
	"fmt"
	documenter "github.com/dl-nft-books/blob-svc/connector/api"
	"github.com/dl-nft-books/core-svc/internal/config"
	"github.com/dl-nft-books/core-svc/internal/data"
	"github.com/dl-nft-books/core-svc/internal/data/postgres"
	"github.com/dl-nft-books/core-svc/resources"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/running"
	"net/http"
)

var (
	documentDeleteError = errors.New("failed to delete document from S3")
)

type TaskCleaner struct {
	name       string
	logger     *logan.Entry
	db         data.DB
	selector   data.TaskSelector
	cleanerCfg config.TaskCleanerCfg
	documenter *documenter.Connector
}

func New(cfg config.Config) *TaskCleaner {
	status := resources.TaskFinishedGeneration
	period := cfg.TaskCleanerCfg().CleaningPeriod

	return &TaskCleaner{
		name:   cfg.TaskCleanerCfg().Name,
		db:     postgres.NewDB(cfg.DB()),
		logger: cfg.Log(),

		selector: data.TaskSelector{
			Status: &status,
			Period: &period,
		},
		cleanerCfg: cfg.TaskCleanerCfg(),
		documenter: cfg.DocumenterConnector(),
	}
}

func (p *TaskCleaner) Run(ctx context.Context) {
	running.WithBackOff(
		ctx, p.logger,
		p.name, p.run,
		p.cleanerCfg.Runner.NormalPeriod,
		p.cleanerCfg.Runner.MinAbnormalPeriod,
		p.cleanerCfg.Runner.MaxAbnormalPeriod,
	)
}

func (p *TaskCleaner) run(ctx context.Context) error {
	unresolvedTasks, err := p.getUnresolvedTasks()
	if err != nil {
		return errors.Wrap(err, "failed to get tasks from the database")
	}

	if len(unresolvedTasks) == 0 {
		p.logger.Debug("Found no unresolved tasks to process")
		return nil
	}

	for _, task := range unresolvedTasks {
		errFields := logan.F{
			"task_id":     task.Id,
			"task_status": task.Status,
		}

		bannerName := fmt.Sprintf("%s.png", task.BannerIpfsHash)

		statusCode, err := p.documenter.DeleteDocument(bannerName)

		if err != nil {
			return errors.Wrap(err, "failed to delete document from S3", errFields)
		}

		if statusCode != http.StatusOK {
			p.logger.WithFields(logan.F{"status_code": statusCode}).Error(documentDeleteError)
			return documentDeleteError
		}

		if err = p.db.New().Tasks().Delete(task.Id); err != nil {
			return errors.Wrap(err, "failed to delete task from data base", errFields)
		}

		p.logger.Debugf("Document deleted from S3 (task_id: %d)", task.Id)

	}

	return nil
}

func (p *TaskCleaner) getUnresolvedTasks() ([]data.Task, error) {
	return p.db.New().Tasks().Select(p.selector)
}
