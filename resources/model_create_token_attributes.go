/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type CreateTokenAttributes struct {
	// Address of a user who purchased this token
	Account string `json:"account"`
	// network chain id
	ChainId int64 `json:"chain_id"`
	// true if user payed by token and false if by nft exchange
	IsTokenPayment bool `json:"is_token_payment"`
	// Hash of a metadata file
	MetadataHash string `json:"metadata_hash"`
	// personal string on first page
	Signature string `json:"signature"`
	// status of uploading
	Status TokenStatus `json:"status"`
	// id of token
	TokenId int64 `json:"token_id"`
}
