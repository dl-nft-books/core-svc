package config

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

const yamlIpfsKey = "ipfs"

type Ipfs interface {
	BaseUri() string
}

type ipfs struct {
	getter kv.Getter
	once   comfig.Once
}

func NewIpfs(getter kv.Getter) Ipfs {
	return &ipfs{
		getter: getter,
	}
}

type IpfsCfg struct {
	BaseUri string `fig:"base_uri,required"`
}

func (i *ipfs) BaseUri() string {
	return i.once.Do(func() interface{} {
		var cfg IpfsCfg

		err := figure.Out(&cfg).
			From(kv.MustGetStringMap(i.getter, yamlIpfsKey)).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out ipfs fields"))
		}
		return cfg.BaseUri
	}).(string)
}
