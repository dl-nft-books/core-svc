package signature

import "math/big"

type SignatureParameters struct {
	R string `json:"r"`
	S string `json:"s"`
	V int    `json:"v"`
}

type EIP712DomainData struct {
	ContractName     string
	ContractVersion  string
	VerifyingAddress string
	ChainID          int64
}

type MintInfo struct {
	TokenAddress     string
	TokenURI         string
	PricePerOneToken *big.Int
	EndTimestamp     int64

	HashedTokenURI []byte
}

type CreateInfo struct {
	TokenContractId  int64
	TokenName        string
	TokenSymbol      string
	PricePerOneToken *big.Int

	HashedTokenName   []byte
	HashedTokenSymbol []byte
}
