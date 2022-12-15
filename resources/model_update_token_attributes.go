/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type UpdateTokenAttributes struct {
	// Token's metadata hash
	MetadataHash *string `json:"metadata_hash,omitempty"`
	// Address of a user who purchased this token
	Owner   *string      `json:"owner,omitempty"`
	Status  *TokenStatus `json:"status,omitempty"`
	TokenId *int64       `json:"token_id,omitempty"`
}
