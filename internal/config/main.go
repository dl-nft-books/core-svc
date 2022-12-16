package config

import (
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/copus"
	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/kit/pgdb"

	s3config "gitlab.com/tokend/nft-books/blob-svc/connector/config"
	booker "gitlab.com/tokend/nft-books/book-svc/connector"
	tracker "gitlab.com/tokend/nft-books/contract-tracker/connector"
	doormanCfg "gitlab.com/tokend/nft-books/doorman/connector/config"
	pricer "gitlab.com/tokend/nft-books/price-svc/connector"
)

type Config interface {
	// Default config interfaces
	comfig.Logger
	types.Copuser
	comfig.Listenerer
	pgdb.Databaser

	// Connectors
	booker.Booker
	pricer.Pricer
	s3config.Documenter
	tracker.Tracker
	doormanCfg.DoormanConfiger

	// Internal service configuration
	MintConfigurator
	TaskProcessor
	TaskCleaner
	PdfSignatureParams() *SignatureParams
	ApiRestrictions() ApiRestrictions
}
type config struct {
	// Default config interfaces
	comfig.Logger
	types.Copuser
	comfig.Listenerer
	pgdb.Databaser

	// Connectors
	booker.Booker
	pricer.Pricer
	s3config.Documenter
	tracker.Tracker
	doormanCfg.DoormanConfiger

	// Internal service configuration
	MintConfigurator
	TaskProcessor
	TaskCleaner
	pdfSignatureParams comfig.Once

	// Getters and comfig.Once's
	getter  kv.Getter
	apiOnce comfig.Once
}

func New(getter kv.Getter) Config {
	return &config{
		// Default config interfaces
		getter:     getter,
		Copuser:    copus.NewCopuser(getter),
		Listenerer: comfig.NewListenerer(getter),
		Databaser:  pgdb.NewDatabaser(getter),
		Logger:     comfig.NewLogger(getter, comfig.LoggerOpts{}),

		// Connectors
		Documenter:      s3config.NewDocumenter(getter),
		Pricer:          pricer.NewPricer(getter),
		Booker:          booker.NewBooker(getter),
		Tracker:         tracker.NewTracker(getter),
		DoormanConfiger: doormanCfg.NewDoormanConfiger(getter),

		// Internal service configuration
		MintConfigurator: NewEthMinterConfigurator(getter),
		TaskProcessor:    NewTaskProcessor(getter),
		TaskCleaner:      NewTaskCleaner(getter),
	}
}
