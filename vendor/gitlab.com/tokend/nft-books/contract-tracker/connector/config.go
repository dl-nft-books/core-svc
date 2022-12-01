package connector

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

const trackerYamlKey = "trackerYamlKey"

type Tracker interface {
	TrackerConnector() *Connector
}

type tracker struct {
	once   comfig.Once
	getter kv.Getter
}

type trackerCfg struct {
	URL   string `fig:"url,required"`
	Token string `fig:"token,required"`
}

func NewTracker(getter kv.Getter) Tracker {
	return &tracker{getter: getter}
}

func (c *tracker) TrackerConnector() *Connector {
	return c.once.Do(func() interface{} {
		var cfg trackerCfg

		raw := kv.MustGetStringMap(c.getter, trackerYamlKey)

		if err := figure.
			Out(&cfg).
			From(raw).
			Please(); err != nil {
			panic(errors.Wrap(err, "failed to figure out"))
		}

		return NewConnector(cfg.Token, cfg.URL)
	}).(*Connector)
}
