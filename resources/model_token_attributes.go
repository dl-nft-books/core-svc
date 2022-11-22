/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type TokenAttributes struct {
	// Token's description retrieved from json metadata
	Description string `json:"description"`
	// Url to the token's image
	ImageUrl string `json:"image_url"`
	// Token's name retrieved from json metadata
	Name string `json:"name"`
	// Address of a user who purchased this token
	Owner     string      `json:"owner"`
	Signature string      `json:"signature"`
	Status    TokenStatus `json:"status"`
	TokenId   int32       `json:"token_id"`
}
