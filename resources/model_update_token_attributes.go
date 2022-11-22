/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type UpdateTokenAttributes struct {
	// Address of a user who purchased this token
	Owner   *string      `json:"owner,omitempty"`
	Status  *TokenStatus `json:"status,omitempty"`
	TokenId *int32       `json:"token_id,omitempty"`
}
