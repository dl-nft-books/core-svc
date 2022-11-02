package config

import (
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/copus"
	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/kit/kv"
	doormanCfg "gitlab.com/tokend/nft-books/doorman/connector/config"
)

type Config interface {
	comfig.Logger
	types.Copuser
	comfig.Listenerer
	MimeTypesConfigurator
	AWSConfigurator

	doormanCfg.DoormanConfiger
}

type config struct {
	comfig.Logger
	types.Copuser
	comfig.Listenerer
	getter kv.Getter
	MimeTypesConfigurator
	AWSConfigurator

	doormanCfg.DoormanConfiger
}

func New(getter kv.Getter) Config {
	return &config{
		getter:                getter,
		Copuser:               copus.NewCopuser(getter),
		Listenerer:            comfig.NewListenerer(getter),
		Logger:                comfig.NewLogger(getter, comfig.LoggerOpts{}),
		MimeTypesConfigurator: NewMimeTypesConfigurator(getter),
		AWSConfigurator:       NewAWSConfigurator(getter),
		DoormanConfiger:       doormanCfg.NewDoormanConfiger(getter),
	}
}
