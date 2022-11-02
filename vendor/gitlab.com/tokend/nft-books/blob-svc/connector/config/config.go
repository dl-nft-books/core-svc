package config

import (
	"net/http"
	"net/url"

	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/connectors/signed"
	"gitlab.com/tokend/nft-books/blob-svc/connector/api"
	"gitlab.com/tokend/nft-books/blob-svc/internal/config"
)

type Documenter interface {
	DocumenterConnector() *api.Connector
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

func (c *documenter) DocumenterConnector() *api.Connector {
	return c.once.Do(func() interface{} {
		var conf documenterConfig

		var awsConf config.AWSConfig

		if err := figure.
			Out(&conf).
			With(figure.BaseHooks).
			From(kv.MustGetStringMap(c.getter, "documenter")).
			Please(); err != nil {
			panic(errors.Wrap(err, "failed to figure out documenter"))
		}

		if err := figure.
			Out(&awsConf).
			From(kv.MustGetStringMap(c.getter, "aws")).
			Please(); err != nil {
			panic(errors.Wrap(err, "failed to figure out aws config"))
		}

		cli := signed.NewClient(http.DefaultClient, conf.URL)

		return api.NewConnector(cli, awsConf, conf.Token)
	}).(*api.Connector)
}
