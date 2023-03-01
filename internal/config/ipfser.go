package config

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"time"
)

const yamlIpfserKey = "ipfser"

type Ipfser interface {
	IpfserCfg() IpfserCfg
}

type ipfser struct {
	getter kv.Getter
	once   comfig.Once
}

func NewIpfser(getter kv.Getter) Ipfser {
	return &ipfser{
		getter: getter,
	}
}

type IpfserCfg struct {
	BaseUri         string        `fig:"base_uri,required"`
	NumberOfRetries int           `fig:"number_of_retries,required"`
	RetryPeriod     time.Duration `fig:"retry_period,required"`
}

func (i *ipfser) IpfserCfg() IpfserCfg {
	return i.once.Do(func() interface{} {
		var cfg IpfserCfg

		err := figure.Out(&cfg).
			From(kv.MustGetStringMap(i.getter, yamlIpfserKey)).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out ipfser fields"))
		}
		return cfg
	}).(IpfserCfg)
}
