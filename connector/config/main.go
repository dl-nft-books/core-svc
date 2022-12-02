package config

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/generator-svc/connector/api"
)

type GeneratorConfigurator interface {
	GeneratorConnector() *api.Connector
}

type generatorConfigurator struct {
	once   comfig.Once
	getter kv.Getter
}

type GeneratorConnectorConfig struct {
	URL   string `fig:"url,required"`
	Token string `fig:"token,required"`
}

func NewGeneratorConfigurator(getter kv.Getter) GeneratorConfigurator {
	return &generatorConfigurator{getter: getter}
}

func (c *generatorConfigurator) GeneratorConnector() *api.Connector {
	return c.once.Do(func() interface{} {
		config := GeneratorConnectorConfig{}

		raw := kv.MustGetStringMap(c.getter, "connector")

		if err := figure.
			Out(&config).
			From(raw).
			Please(); err != nil {
			panic(errors.Wrap(err, "failed to figure out"))
		}

		return api.NewConnector(config.Token, config.URL)
	}).(*api.Connector)
}
