/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type TokenAttributes struct {
	// Token's description retrieved from json metadata
	Description string `json:"description"`
	// Url to the token's image
	ImageUrl string `json:"image_url"`
	// true if user payed by token and false if by nft exchange
	IsTokenPayment bool `json:"is_token_payment"`
	// Hash of a metadata file
	MetadataHash string `json:"metadata_hash"`
	// Token's name retrieved from json metadata
	Name string `json:"name"`
	// Address of a user who purchased this token
	Owner string `json:"owner"`
	// personal string on first page
	Signature string `json:"signature"`
	// token uploading status
	Status TokenStatus `json:"status"`
	// Id of token
	TokenId int64 `json:"token_id"`
}
