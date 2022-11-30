package config

import (
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/copus"
	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/kit/kv"
	s3connector "gitlab.com/tokend/nft-books/blob-svc/connector/api"
	s3config "gitlab.com/tokend/nft-books/blob-svc/connector/config"
	booksConnector "gitlab.com/tokend/nft-books/book-svc/connector/config"
	networkerCfg "gitlab.com/tokend/nft-books/network-svc/connector/config"
	pricer "gitlab.com/tokend/nft-books/price-svc/connector"
)

type Config interface {
	comfig.Logger
	types.Copuser
	comfig.Listenerer
	Databaser
	EthMinterConfigurator
	TaskProcessor

	networkerCfg.NetworkConfigurator

	booksConnector.BooksConnectorConfigurator
	DocumenterConnector() *s3connector.Connector
	PdfSignatureParams() *SignatureParams
	ApiRestrictions() ApiRestrictions
	pricer.Pricer
}
type config struct {
	comfig.Logger
	types.Copuser
	comfig.Listenerer
	Databaser
	EthMinterConfigurator
	TaskProcessor
	s3config.Documenter
	pdfSignatureParams comfig.Once
	pricer.Pricer
	networkerCfg.NetworkConfigurator
	booksConnector.BooksConnectorConfigurator

	getter  kv.Getter
	apiOnce comfig.Once
}

func New(getter kv.Getter) Config {
	return &config{
		getter:                     getter,
		Copuser:                    copus.NewCopuser(getter),
		Listenerer:                 comfig.NewListenerer(getter),
		Logger:                     comfig.NewLogger(getter, comfig.LoggerOpts{}),
		EthMinterConfigurator:      NewEthMinterConfigurator(getter),
		Databaser:                  NewDatabaser(getter),
		TaskProcessor:              NewTaskProcessor(getter),
		Documenter:                 s3config.NewDocumenter(getter),
		Pricer:                     pricer.NewPricer(getter),
		NetworkConfigurator:        networkerCfg.NewNetworkConfigurator(getter),
		BooksConnectorConfigurator: booksConnector.NewBooksConfigurator(getter),
	}
}
