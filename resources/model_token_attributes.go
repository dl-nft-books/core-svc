/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type TokenAttributes struct {
	ChainId int64 `json:"chain_id"`
	// Token's description retrieved from json metadata
	Description string `json:"description"`
	// Url to the token's image
	ImageUrl string `json:"image_url"`
	// Hash of a metadata file
	MetadataHash string `json:"metadata_hash"`
	// Token's name retrieved from json metadata
	Name string `json:"name"`
	// Address of a user who purchased this token
	Owner     string      `json:"owner"`
	Signature string      `json:"signature"`
	Status    TokenStatus `json:"status"`
	TokenId   int32       `json:"token_id"`
}
