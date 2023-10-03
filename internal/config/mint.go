package config

import (
	"fmt"
	"reflect"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type MintConfigurator interface {
	MintConfig() *MintConfig
}

type MintConfig struct {
	Precision  int           `fig:"precision,required"`
	Expiration time.Duration `fig:"expiration,required"`
}

type ethMinterConfigurator struct {
	getter kv.Getter
	once   comfig.Once
}

func NewEthMinterConfigurator(getter kv.Getter) MintConfigurator {
	return &ethMinterConfigurator{
		getter: getter,
	}
}

func (c *ethMinterConfigurator) MintConfig() *MintConfig {
	return c.once.Do(func() interface{} {
		var cfg MintConfig

		err := figure.Out(&cfg).
			With(figure.BaseHooks, ecdsaHook).
			From(kv.MustGetStringMap(c.getter, "mint")).
			Please()
		if err != nil {
			panic(err)
		}

		return &cfg
	}).(*MintConfig)
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
