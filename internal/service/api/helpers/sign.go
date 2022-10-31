package helpers

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	signer "github.com/ethersphere/bee/pkg/crypto"
	"github.com/ethersphere/bee/pkg/crypto/eip712"
	sha3 "github.com/miguelmota/go-solidity-sha3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/generator-svc/internal/config"
)

const defaultAddress = "0x0000000000000000000000000000000000000000"

type SignatureParameters struct {
	R string `json:"r"`
	S string `json:"s"`
	V int    `json:"v"`
}

type SignInfo struct {
	ContractName    string
	ContractVersion string
	ContractAddress string
	TokenAddress    string
	TokenURI        string
	Price           *big.Int
	ChainID         int64
	EndTimestamp    int64
	HashedTokenURI  []byte
}

func Sign(info *SignInfo, config *config.EthMinterConfig) (*SignatureParameters, error) {
	privateKey := config.PrivateKey

	// hashing token uri
	tokenURIsha3 := sha3.String(info.TokenURI)
	info.HashedTokenURI = sha3.SoliditySHA3(tokenURIsha3)

	// if token address is not specified -- setting default address
	if info.TokenAddress == "" {
		info.TokenAddress = defaultAddress
	}

	signature, err := signEIP712(privateKey, info)
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign EIP712 hash")
	}

	return parseSignatureParameters(signature)
}

func signEIP712(privateKey *ecdsa.PrivateKey, info *SignInfo) ([]byte, error) {
	spew.Dump(info)

	data := &eip712.TypedData{
		Types: apitypes.Types{
			"Mint": []apitypes.Type{
				{Name: "paymentTokenAddress", Type: "address"},
				{Name: "paymentTokenPrice", Type: "uint256"},
				{Name: "endTimestamp", Type: "uint256"},
				{Name: "tokenURI", Type: "bytes32"},
			},
			"EIP712Domain": []apitypes.Type{
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
				{Name: "verifyingContract", Type: "address"},
			},
		},
		PrimaryType: "Mint",
		Domain: apitypes.TypedDataDomain{
			Name:              info.ContractName,
			Version:           info.ContractVersion,
			ChainId:           math.NewHexOrDecimal256(info.ChainID),
			VerifyingContract: info.ContractAddress,
		},
		Message: apitypes.TypedDataMessage{
			"paymentTokenAddress": info.TokenAddress,
			"paymentTokenPrice":   info.Price.String(),
			"endTimestamp":        math.NewHexOrDecimal256(info.EndTimestamp),
			"tokenURI":            info.HashedTokenURI,
		},
	}

	return signer.NewDefaultSigner(privateKey).SignTypedData(data)
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
