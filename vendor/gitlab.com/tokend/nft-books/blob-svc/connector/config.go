package connector

import (
	"net/http"
	"net/url"

	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/connectors/signed"
)

type Documenter interface {
	DocumenterConnector() *Connector
}

type documenter struct {
	once   comfig.Once
	getter kv.Getter
}

func NewDocumenter(getter kv.Getter) Documenter {
	return &documenter{
		getter: getter,
	}
}

type documenterConfig struct {
	URL   *url.URL `fig:"url,required"`
	Token string   `fig:"token,required"`
}

func (c *documenter) DocumenterConnector() *Connector {
	return c.once.Do(func() interface{} {
		var config documenterConfig

		err := figure.
			Out(&config).
			With(figure.BaseHooks).
			From(kv.MustGetStringMap(c.getter, "documenter")).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out documenter"))
		}

		cli := signed.NewClient(http.DefaultClient, config.URL)

		return NewConnector(cli, config.Token)
	}).(*Connector)
}
