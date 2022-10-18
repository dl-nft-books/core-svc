package config

import (
	"crypto/ecdsa"
	"fmt"
	"reflect"

	"github.com/ethereum/go-ethereum/crypto"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type EthMinterConfig struct {
	PrivateKey *ecdsa.PrivateKey `fig:"eth_signer,required"`
	Precision  int               `fig:"precision,required"`
}

var hooks = figure.Hooks{
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

func (c *config) EthMinter() EthMinterConfig {
	return c.ethMinter.Do(func() interface{} {
		var cfg EthMinterConfig

		err := figure.Out(&cfg).
			With(figure.BaseHooks, hooks).
			From(kv.MustGetStringMap(c.getter, "eth_minter")).
			Please()
		if err != nil {
			panic(err)
		}

		return cfg
	}).(EthMinterConfig)
}
