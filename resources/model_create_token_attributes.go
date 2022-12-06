/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type CreateTokenAttributes struct {
	// Address of a user who purchased this token
	Account string `json:"account"`
	// Hash of a metadata file
	MetadataHash string      `json:"metadata_hash"`
	Signature    string      `json:"signature"`
	Status       TokenStatus `json:"status"`
	TokenId      int64       `json:"token_id"`
}
