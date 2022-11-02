package config

import (
	"time"

	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
)

type AWSConfigurator interface {
	AWSConfig() *AWSConfig
}

type AWSConfig struct {
	Endpoint       string        `fig:"endpoint"`
	AccessKeyID    string        `fig:"access_key,required"`
	SecretKeyID    string        `fig:"secret_key,required"`
	Bucket         string        `fig:"bucket,required"`
	Expiration     time.Duration `fig:"expiration,required"`
	SslDisable     bool          `fig:"ssldisable,required"`
	ForcePathStyle bool          `fig:"force_path_style,required"`
	Region         string        `fig:"region,required"`
}

func NewAWSConfigurator(getter kv.Getter) AWSConfigurator {
	return &awsConfig{
		getter: getter,
	}
}

type awsConfig struct {
	getter kv.Getter
	once   comfig.Once
}

func (c *awsConfig) AWSConfig() *AWSConfig {
	return c.once.Do(func() interface{} {
		config := AWSConfig{}

		if err := figure.Out(&config).From(kv.MustGetStringMap(c.getter, "aws")).Please(); err != nil {
			panic(err)
		}

		return &config
	}).(*AWSConfig)
}
