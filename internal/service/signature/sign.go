package signature

import (
	"crypto/ecdsa"

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

func SignMintInfo(
	mintInfo *MintInfo,
	domainData *EIP712DomainData,
	config *config.EthMinterConfig,
) (
	*SignatureParameters,
	error,
) {
	privateKey := config.PrivateKey

	// hashing token uri
	tokenURIRaw := sha3.String(mintInfo.TokenURI)
	mintInfo.HashedTokenURI = sha3.SoliditySHA3(tokenURIRaw)

	// if token address is not specified -- setting default address
	if mintInfo.TokenAddress == "" {
		mintInfo.TokenAddress = defaultAddress
	}

	signature, err := signMintInfoByEIP712(privateKey, mintInfo, domainData)
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign EIP712 hash")
	}

	return parseSignatureParameters(signature)
}

func SignCreateInfo(
	createInfo *CreateInfo,
	domainData *EIP712DomainData,
	config *config.EthMinterConfig,
) (
	*SignatureParameters,
	error,
) {
	privateKey := config.PrivateKey

	// hashing token params
	tokenNameRaw := sha3.String(createInfo.TokenName)
	createInfo.HashedTokenName = sha3.SoliditySHA3(tokenNameRaw)

	tokenSymbolRaw := sha3.String(createInfo.TokenSymbol)
	createInfo.HashedTokenSymbol = sha3.SoliditySHA3(tokenSymbolRaw)

	signature, err := signCreateInfoByEIP712(privateKey, createInfo, domainData)
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign EIP712 hash")
	}

	return parseSignatureParameters(signature)
}

func signCreateInfoByEIP712(
	privateKey *ecdsa.PrivateKey,
	createInfo *CreateInfo,
	domainData *EIP712DomainData,
) (
	[]byte,
	error,
) {
	spew.Dump(createInfo)

	data := &eip712.TypedData{
		Types: apitypes.Types{
			"Create": []apitypes.Type{
				{Name: "tokenContractId", Type: "uint256"},
				{Name: "tokenName", Type: "bytes32"},
				{Name: "tokenSymbol", Type: "bytes32"},
				{Name: "pricePerOneToken", Type: "uint256"},
			},
			"EIP712Domain": []apitypes.Type{
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
				{Name: "verifyingContract", Type: "address"},
			},
		},
		PrimaryType: "Create",
		Domain: apitypes.TypedDataDomain{
			Name:              domainData.ContractName,
			Version:           domainData.ContractVersion,
			ChainId:           math.NewHexOrDecimal256(domainData.ChainID),
			VerifyingContract: domainData.VerifyingAddress,
		},
		Message: apitypes.TypedDataMessage{
			"tokenContractId":  math.NewHexOrDecimal256(createInfo.TokenContractId),
			"tokenName":        createInfo.HashedTokenName,
			"tokenSymbol":      createInfo.HashedTokenSymbol,
			"pricePerOneToken": createInfo.PricePerOneToken.String(),
		},
	}

	return signer.NewDefaultSigner(privateKey).SignTypedData(data)
}

func signMintInfoByEIP712(privateKey *ecdsa.PrivateKey,
	mintInfo *MintInfo,
	domainData *EIP712DomainData,
) (
	[]byte,
	error,
) {
	spew.Dump(mintInfo)

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
			Name:              domainData.ContractName,
			Version:           domainData.ContractVersion,
			ChainId:           math.NewHexOrDecimal256(domainData.ChainID),
			VerifyingContract: domainData.VerifyingAddress,
		},
		Message: apitypes.TypedDataMessage{
			"paymentTokenAddress": mintInfo.TokenAddress,
			"paymentTokenPrice":   mintInfo.PricePerOneToken.String(),
			"endTimestamp":        math.NewHexOrDecimal256(mintInfo.EndTimestamp),
			"tokenURI":            mintInfo.HashedTokenURI,
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
