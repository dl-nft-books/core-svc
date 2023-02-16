package signature

import (
	"crypto/ecdsa"
	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	signer "github.com/ethersphere/bee/pkg/crypto"
	"github.com/ethersphere/bee/pkg/crypto/eip712"
	sha3 "github.com/miguelmota/go-solidity-sha3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/generator-svc/internal/config"
	"math/big"
)

func SignMintByNftInfo(
	mintInfo *MintByNftInfo,
	domainData *EIP712DomainData,
	config *config.MintConfig,
) (
	*Parameters,
	error,
) {
	privateKey := config.PrivateKey
	// hashing token uri
	tokenURIRaw := sha3.String(mintInfo.TokenURI)
	mintInfo.HashedTokenURI = sha3.SoliditySHA3(tokenURIRaw)

	spew.Dump(*mintInfo)
	spew.Dump(*domainData)
	signature, err := signMintByNftInfoByEIP712(privateKey, mintInfo, domainData)
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign EIP712 hash")
	}

	return parseSignatureParameters(signature)
}

func signMintByNftInfoByEIP712(privateKey *ecdsa.PrivateKey,
	mintInfo *MintByNftInfo,
	domainData *EIP712DomainData,
) (
	[]byte,
	error,
) {
	data := &eip712.TypedData{
		Types: apitypes.Types{
			"Mint": []apitypes.Type{
				{Name: "paymentTokenAddress", Type: "address"},
				{Name: "paymentTokenPrice", Type: "uint256"},
				{Name: "discount", Type: "uint256"},
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
			"paymentTokenAddress": mintInfo.NftAddress,
			"paymentTokenPrice":   mintInfo.NftFloorPrice.String(),
			"discount":            big.NewInt(0).String(),
			"endTimestamp":        math.NewHexOrDecimal256(mintInfo.EndTimestamp),
			"tokenURI":            mintInfo.HashedTokenURI,
		},
	}

	return signer.NewDefaultSigner(privateKey).SignTypedData(data)
}
