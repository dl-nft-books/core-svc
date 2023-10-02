package config

import (
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/copus"
	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/kit/pgdb"

	s3config "github.com/dl-nft-books/blob-svc/connector/config"
	booker "github.com/dl-nft-books/book-svc/connector"
	doormanCfg "github.com/dl-nft-books/doorman/connector/config"
	networker "github.com/dl-nft-books/network-svc/connector"
	pricer "github.com/dl-nft-books/price-svc/connector"
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
	doormanCfg.DoormanConfiger
	networker.NetworkConfigurator

	// Internal service configuration
	MintConfigurator
	AcceptConfigurator
	TransactionConfigurator
	Promocoder
	TaskCleaner
	Ipfser
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
	doormanCfg.DoormanConfiger
	networker.NetworkConfigurator

	// Internal service configuration
	MintConfigurator
	AcceptConfigurator
	TransactionConfigurator
	Promocoder
	TaskCleaner
	Ipfser

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
		Documenter:          s3config.NewDocumenter(getter),
		Pricer:              pricer.NewPricer(getter),
		Booker:              booker.NewBooker(getter),
		NetworkConfigurator: networker.NewNetworkConfigurator(getter),
		DoormanConfiger:     doormanCfg.NewDoormanConfiger(getter),

		// Internal service configuration
		MintConfigurator:        NewEthMinterConfigurator(getter),
		AcceptConfigurator:      NewEthAccepterConfigurator(getter),
		TransactionConfigurator: NewEthTransactionerConfigurator(getter),
		Promocoder:              NewPromocoder(getter),
		TaskCleaner:             NewTaskCleaner(getter),
		Ipfser:                  NewIpfser(getter),
	}
}
