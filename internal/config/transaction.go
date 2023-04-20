package config

import (
	"crypto/ecdsa"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
)

type TransactionConfigurator interface {
	TransactionConfig() *TransactionConfig
}

type TransactionConfig struct {
	PrivateKey *ecdsa.PrivateKey `fig:"private_key,required"`
	GasLimit   uint64            `fig:"gas_limit,required"`
}

type ethTransactionerConfigurator struct {
	getter kv.Getter
	once   comfig.Once
}

func NewEthTransactionerConfigurator(getter kv.Getter) TransactionConfigurator {
	return &ethTransactionerConfigurator{
		getter: getter,
	}
}

func (c *ethTransactionerConfigurator) TransactionConfig() *TransactionConfig {
	return c.once.Do(func() interface{} {
		var cfg TransactionConfig

		err := figure.Out(&cfg).
			With(figure.BaseHooks, ecdsaHook).
			From(kv.MustGetStringMap(c.getter, "transaction")).
			Please()
		if err != nil {
			panic(err)
		}

		return &cfg
	}).(*TransactionConfig)
}
