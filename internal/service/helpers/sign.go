package helpers

import (
	"crypto/ecdsa"
	"fmt"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type SignatureParameters struct {
	R string `json:"r"`
	S string `json:"s"`
	V int    `json:"v"`
}

func Sign(r *http.Request, price, contractAddress string) (*SignatureParameters, error) {
	mintConfig := Minter(r)
	privateKey := mintConfig.PrivateKey

	priceToSign := fmt.Sprintf("%s%s", price, strings.Repeat("0", mintConfig.Precision))

	signature, err := signEIP712(privateKey, priceToSign, contractAddress)
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign EIP712 hash")
	}
	return parseSignatureParameters(signature)
}

func signEIP712(privateKey *ecdsa.PrivateKey, price, contractAddress string) ([]byte, error) {

}

func parseSignatureParameters(signature []byte) (*SignatureParameters, error) {
	if len(signature) != 65 {
		return nil, errors.New("bad signature")
	}
	params := SignatureParameters{}

	params.R = hexutil.Encode(signature[:32])
	params.S = hexutil.Encode(signature[32:64])
	params.V = int(signature[64])

	return &params, nil
}
