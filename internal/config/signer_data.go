package config

import (
	"crypto/ecdsa"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
)

type SignerDataConfigurator interface {
	SignerDataConfig() *SignerDataConfig
}

type SignerDataConfig struct {
	PrivateKey *ecdsa.PrivateKey `fig:"private_key,required"`
}

type signerDataConfigurator struct {
	getter kv.Getter
	once   comfig.Once
}

func NewEthSignererDataConfigurator(getter kv.Getter) SignerDataConfigurator {
	return &signerDataConfigurator{
		getter: getter,
	}
}

func (c *signerDataConfigurator) SignerDataConfig() *SignerDataConfig {
	return c.once.Do(func() interface{} {
		var cfg SignerDataConfig

		err := figure.Out(&cfg).
			With(figure.BaseHooks, ecdsaHook).
			From(kv.MustGetStringMap(c.getter, "signer_data")).
			Please()
		if err != nil {
			panic(err)
		}

		return &cfg
	}).(*SignerDataConfig)
}
