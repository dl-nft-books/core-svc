package config

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"time"
)

const (
	apiYamlkey = "api"
)

type ApiRestrictions struct {
	RequestDelay      time.Duration `fig:"request_delay"`
	MaxFailedAttempts uint64        `fig:"max_failed_attempts"`
}

var defaultApiRestrictions = ApiRestrictions{
	RequestDelay:      5 * time.Minute,
	MaxFailedAttempts: 10,
}

func (c *config) ApiRestrictions() ApiRestrictions {
	return c.apiOnce.Do(func() interface{} {
		cfg := defaultApiRestrictions

		err := figure.Out(&cfg).
			With(figure.BaseHooks).
			From(kv.MustGetStringMap(c.getter, apiYamlkey)).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out api settings from config"))
		}

		return cfg
	}).(ApiRestrictions)
}
