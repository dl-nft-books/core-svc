package config

import (
	"crypto/ecdsa"
	"time"

	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
)

type AcceptConfigurator interface {
	AcceptConfig() *AcceptConfig
}

type AcceptConfig struct {
	PrivateKey *ecdsa.PrivateKey `fig:"signer,required"`
	Precision  int               `fig:"precision,required"`
	Expiration time.Duration     `fig:"expiration,required"`
}

type ethAccepterConfigurator struct {
	getter kv.Getter
	once   comfig.Once
}

func NewEthAccepterConfigurator(getter kv.Getter) AcceptConfigurator {
	return &ethAccepterConfigurator{
		getter: getter,
	}
}

func (c *ethAccepterConfigurator) AcceptConfig() *AcceptConfig {
	return c.once.Do(func() interface{} {
		var cfg AcceptConfig

		err := figure.Out(&cfg).
			With(figure.BaseHooks, ecdsaHook).
			From(kv.MustGetStringMap(c.getter, "accept")).
			Please()
		if err != nil {
			panic(err)
		}

		return &cfg
	}).(*AcceptConfig)
}
