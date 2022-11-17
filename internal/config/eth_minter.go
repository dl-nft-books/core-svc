package config

import (
	"crypto/ecdsa"
	"fmt"
	"reflect"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type EthMinterConfigurator interface {
	EthMinter() *EthMinterConfig
}

type EthMinterConfig struct {
	PrivateKey *ecdsa.PrivateKey `fig:"eth_signer,required"`
	ChainID    int64             `fig:"chain_id,required"`
	Precision  int               `fig:"precision,required"`
	Expiration time.Duration     `fig:"expiration,required"`

	TokenFactoryName    string `fig:"token_factory_name,required"`
	TokenFactoryAddress string `fig:"token_factory_address,required"`
	TokenFactoryVersion string `fig:"token_factory_version,required"`
}

type ethMinterConfigurator struct {
	getter kv.Getter
	once   comfig.Once
}

func NewEthMinterConfigurator(getter kv.Getter) EthMinterConfigurator {
	return &ethMinterConfigurator{
		getter: getter,
	}
}

func (c *ethMinterConfigurator) EthMinter() *EthMinterConfig {
	return c.once.Do(func() interface{} {
		var cfg EthMinterConfig

		err := figure.Out(&cfg).
			With(figure.BaseHooks, ecdsaHook).
			From(kv.MustGetStringMap(c.getter, "eth_minter")).
			Please()
		if err != nil {
			panic(err)
		}

		return &cfg
	}).(*EthMinterConfig)
}

var ecdsaHook = figure.Hooks{
	"*ecdsa.PrivateKey": func(value interface{}) (reflect.Value, error) {
		switch v := value.(type) {
		case string:
			privKey, err := crypto.HexToECDSA(v)
			if err != nil {
				return reflect.Value{}, errors.Wrap(err, "invalid hex private key")
			}
			return reflect.ValueOf(privKey), nil
		default:
			return reflect.Value{}, fmt.Errorf("unsupported conversion from %T", value)
		}
	},
}
