package signature

import (
	"crypto/ecdsa"
	"gitlab.com/distributed_lab/logan/v3"

	"github.com/dl-nft-books/core-svc/internal/config"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	signer "github.com/ethersphere/bee/pkg/crypto"
	"github.com/ethersphere/bee/pkg/crypto/eip712"
	sha3 "github.com/miguelmota/go-solidity-sha3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func SignAcceptInfo(
	acceptInfo *AcceptInfo,
	domainData *EIP712DomainData,
	config *config.AcceptConfig,
) (
	[]byte,
	error,
) {
	privateKey := config.PrivateKey

	// hashing token uri
	tokenURIRaw := sha3.String(acceptInfo.TokenURI)
	acceptInfo.HashedTokenURI = sha3.SoliditySHA3(tokenURIRaw)

	signature, err := signAcceptInfoByEIP712(privateKey, acceptInfo, domainData)
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign EIP712 hash")
	}
	return signature, nil
}

func signAcceptInfoByEIP712(privateKey *ecdsa.PrivateKey,
	acceptInfo *AcceptInfo,
	domainData *EIP712DomainData,
) (
	[]byte,
	error,
) {
	data := &eip712.TypedData{
		Types: apitypes.Types{
			"BuyWithRequest": []apitypes.Type{
				{Name: "tokenRecipient", Type: "address"},
				{Name: "requestId", Type: "uint256"},
				{Name: "futureTokenId", Type: "uint256"},
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
		PrimaryType: "BuyWithRequest",
		Domain: apitypes.TypedDataDomain{
			Name:              domainData.ContractName,
			Version:           domainData.ContractVersion,
			ChainId:           math.NewHexOrDecimal256(domainData.ChainID),
			VerifyingContract: domainData.VerifyingAddress,
		},
		Message: apitypes.TypedDataMessage{
			"tokenRecipient": acceptInfo.TokenRecipient,
			"requestId":      math.NewHexOrDecimal256(acceptInfo.RequestId),
			"futureTokenId":  math.NewHexOrDecimal256(acceptInfo.TokenId),
			"endTimestamp":   math.NewHexOrDecimal256(acceptInfo.EndTimestamp),
			"tokenURI":       acceptInfo.HashedTokenURI,
		},
	}

	return signer.NewDefaultSigner(privateKey).SignTypedData(data)
}

func ParseSignatureParameters(signature []byte) (*Parameters, error) {
	if len(signature) != 65 {
		return nil, errors.From(wrongSignatureLengthErr, logan.F{
			"signature": string(signature),
		})
	}

	params := Parameters{}

	params.R = hexutil.Encode(signature[:32])
	params.S = hexutil.Encode(signature[32:64])
	params.V = int8(signature[64])

	return &params, nil
}
