package config

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"time"
)

const yamlTaskCleanerKey = "task_cleaner"

type TaskCleaner interface {
	TaskCleanerCfg() TaskCleanerCfg
}

type taskCleaner struct {
	getter kv.Getter
	once   comfig.Once
}

func NewTaskCleaner(getter kv.Getter) TaskCleaner {
	return &taskCleaner{
		getter: getter,
	}
}

type TaskCleanerCfg struct {
	Name           string        `fig:"name"`
	Runner         RunnerData    `fig:"runner,required"`
	CleaningPeriod time.Duration `fig:"cleaning_period"`
}

func (t *taskCleaner) TaskCleanerCfg() TaskCleanerCfg {
	return t.once.Do(func() interface{} {
		var cfg TaskCleanerCfg

		err := figure.Out(&cfg).
			From(kv.MustGetStringMap(t.getter, yamlTaskCleanerKey)).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out task cleaner fields"))
		}

		return cfg
	}).(TaskCleanerCfg)
}
