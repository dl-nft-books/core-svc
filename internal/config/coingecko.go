package config

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
)

type CoingeckoConfig struct {
	Host     string `json:"host"`
	Endpoint string `json:"endpoint"`
}

func (c *config) Coingecko() *CoingeckoConfig {
	return c.coingecko.Do(func() interface{} {
		var cfg CoingeckoConfig

		err := figure.Out(&cfg).
			From(kv.MustGetStringMap(c.getter, "coingecko")).
			Please()
		if err != nil {
			panic(err)
		}

		return &cfg
	}).(*CoingeckoConfig)
}
