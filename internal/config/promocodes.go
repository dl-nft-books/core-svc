package config

import (
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
)

const yamlPromocodesKey = "promocodes"

type Promocodes interface {
	PromocodesCfg() PromocodesCfg
}

type promocodes struct {
	getter kv.Getter
	once   comfig.Once
}

func NewPromocodes(getter kv.Getter) Promocodes {
	return &promocodes{
		getter: getter,
	}
}

type PromocodesCfg struct {
	Name    string     `fig:"name"`
	Decimal float64    `fig:"decimal,non_zero"`
	Runner  RunnerData `fig:"runner,required"`
}

func (t *promocodes) PromocodesCfg() PromocodesCfg {
	return t.once.Do(func() interface{} {
		var cfg PromocodesCfg

		err := figure.Out(&cfg).
			From(kv.MustGetStringMap(t.getter, yamlPromocodesKey)).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out task producer fields"))
		}

		return cfg
	}).(PromocodesCfg)
}
