package config

import (
	"time"

	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
)

const yamlTaskProcessorKey = "task_processor"

type RunnerData struct {
	NormalPeriod      time.Duration `fig:"normal_period"`
	MinAbnormalPeriod time.Duration `fig:"min_abnormal_period"`
	MaxAbnormalPeriod time.Duration `fig:"max_abnormal_period"`
}

type CleanerData struct {
	CheckingPeriod time.Duration `fig:"checking_period"`
	CleaningPeriod time.Duration `fig:"cleaning_period"`
	Name           string        `fig:"name"`
}

type TaskProcessor interface {
	TaskProcessorCfg() TaskProcessorCfg
}

type taskProcessor struct {
	getter kv.Getter
	once   comfig.Once
}

func NewTaskProcessor(getter kv.Getter) TaskProcessor {
	return &taskProcessor{
		getter: getter,
	}
}

type TaskProcessorCfg struct {
	Name            string      `fig:"name"`
	Cursor          uint64      `fig:"cursor"`
	Limit           uint64      `fig:"limit,non_zero"`
	ProcessesNumber uint64      `fig:"processes_number,non_zero"`
	Runner          RunnerData  `fig:"runner,required"`
	Cleaner         CleanerData `fig:"cleaner,required"`
}

func (t *taskProcessor) TaskProcessorCfg() TaskProcessorCfg {
	return t.once.Do(func() interface{} {
		var cfg TaskProcessorCfg

		err := figure.Out(&cfg).
			From(kv.MustGetStringMap(t.getter, yamlTaskProcessorKey)).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out task producer fields"))
		}

		return cfg
	}).(TaskProcessorCfg)
}
