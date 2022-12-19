package config

import (
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
)

const yamlPromocodesKey = "promocodes"

type Promocoder interface {
	PromocoderCfg() PromocodesCfg
}

type promocoder struct {
	getter kv.Getter
	once   comfig.Once
}

func NewPromocoder(getter kv.Getter) Promocoder {
	return &promocoder{
		getter: getter,
	}
}

type PromocodesCfg struct {
	Name    string     `fig:"name"`
	Decimal int        `fig:"decimal,non_zero"`
	Runner  RunnerData `fig:"runner,required"`
}

func (t *promocoder) PromocoderCfg() PromocodesCfg {
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
