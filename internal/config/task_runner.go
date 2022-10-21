package config

import (
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"time"
)

const yamlTaskRunnerKey = "task_runner"

type RunnerData struct {
	NormalPeriod      time.Duration `fig:"normal_period"`
	MinAbnormalPeriod time.Duration `fig:"min_abnormal_period"`
	MaxAbnormalPeriod time.Duration `fig:"normal_period"`
}

type TaskRunner interface {
	TaskRunner() TaskRunnerCfg
}

type taskRunner struct {
	getter kv.Getter
	once   comfig.Once
}

func NewTaskRunner(getter kv.Getter) TaskRunner {
	return &taskRunner{
		getter: getter,
	}
}

type TaskRunnerCfg struct {
	Name   string     `fig:"name"`
	Cursor uint64     `fig:"cursor"`
	Limit  uint64     `fig:"limit,non_zero"`
	Runner RunnerData `fig:"runner,required"`
}

func (t *taskRunner) TaskRunner() TaskRunnerCfg {
	return t.once.Do(func() interface{} {
		var cfg TaskRunnerCfg

		err := figure.Out(&cfg).
			From(kv.MustGetStringMap(t.getter, yamlTaskRunnerKey)).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out task producer fields"))
		}

		return cfg
	}).(TaskRunnerCfg)
}
