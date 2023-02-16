package signature

import "math/big"

type Parameters struct {
	R string `json:"r"`
	S string `json:"s"`
	V int8   `json:"v"`
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
	Discount         *big.Int
	EndTimestamp     int64

	HashedTokenURI []byte
}
