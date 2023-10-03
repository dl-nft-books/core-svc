package signature

import (
	"crypto/ecdsa"
	"github.com/dl-nft-books/core-svc/internal/config"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	signer "github.com/ethersphere/bee/pkg/crypto"
	"github.com/ethersphere/bee/pkg/crypto/eip712"
	sha3 "github.com/miguelmota/go-solidity-sha3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

var (
	defaultAddress          = common.Address{}.String()
	wrongSignatureLengthErr = errors.New("length of a signature must be 65")
)

func SignMintInfo(
	mintInfo *MintInfo,
	domainData *EIP712DomainData,
	signerConfig *config.SignerDataConfig,
) (
	[]byte,
	error,
) {
	privateKey := signerConfig.PrivateKey

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
	return signature, nil
}

func signMintInfoByEIP712(privateKey *ecdsa.PrivateKey,
	mintInfo *MintInfo,
	domainData *EIP712DomainData,
) (
	[]byte,
	error,
) {
	data := &eip712.TypedData{
		Types: apitypes.Types{
			"Buy": []apitypes.Type{
				{Name: "tokenRecipient", Type: "address"},
				{Name: "tokenContract", Type: "address"},
				{Name: "futureTokenId", Type: "uint256"},
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
		PrimaryType: "Buy",
		Domain: apitypes.TypedDataDomain{
			Name:              domainData.ContractName,
			Version:           domainData.ContractVersion,
			ChainId:           math.NewHexOrDecimal256(domainData.ChainID),
			VerifyingContract: domainData.VerifyingAddress,
		},
		Message: apitypes.TypedDataMessage{
			"tokenRecipient":      mintInfo.TokenRecipient,
			"tokenContract":       mintInfo.TokenContract,
			"futureTokenId":       math.NewHexOrDecimal256(mintInfo.TokenId),
			"paymentTokenAddress": mintInfo.TokenAddress,
			"paymentTokenPrice":   mintInfo.PricePerOneToken.String(),
			"discount":            mintInfo.Discount.String(),
			"endTimestamp":        math.NewHexOrDecimal256(mintInfo.EndTimestamp),
			"tokenURI":            mintInfo.HashedTokenURI,
		},
	}

	return signer.NewDefaultSigner(privateKey).SignTypedData(data)
}
