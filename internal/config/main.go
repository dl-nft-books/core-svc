package config

import (
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/copus"
	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/kit/kv"
	blobconnector "gitlab.com/tokend/nft-books/blob-svc/connector"
	pricer "gitlab.com/tokend/nft-books/price-svc/connector"
)

type Config interface {
	comfig.Logger
	types.Copuser
	comfig.Listenerer
	Databaser
	EthMinterConfigurator
	TaskProcessor

	DocumenterConnector() *blobconnector.Connector
	PdfSignatureParams() *SignatureParams
	pricer.Pricer
}
type config struct {
	comfig.Logger
	types.Copuser
	comfig.Listenerer
	getter kv.Getter
	Databaser
	EthMinterConfigurator
	TaskProcessor
	blobconnector.Documenter
	pdfSignatureParams comfig.Once
	pricer.Pricer
}

func New(getter kv.Getter) Config {
	return &config{
		getter:                getter,
		Copuser:               copus.NewCopuser(getter),
		Listenerer:            comfig.NewListenerer(getter),
		Logger:                comfig.NewLogger(getter, comfig.LoggerOpts{}),
		EthMinterConfigurator: NewEthMinterConfigurator(getter),
		Databaser:             NewDatabaser(getter),
		TaskProcessor:         NewTaskProcessor(getter),
		Documenter:            blobconnector.NewDocumenter(getter),
		Pricer:                pricer.NewPricer(getter),
	}
}
