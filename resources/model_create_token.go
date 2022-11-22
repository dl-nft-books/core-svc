/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type CreateToken struct {
	Key
	Attributes    CreateTokenAttributes    `json:"attributes"`
	Relationships CreateTokenRelationships `json:"relationships"`
}
type CreateTokenRequest struct {
	Data     CreateToken `json:"data"`
	Included Included    `json:"included"`
}

type CreateTokenListRequest struct {
	Data     []CreateToken `json:"data"`
	Included Included      `json:"included"`
	Links    *Links        `json:"links"`
}

// MustCreateToken - returns CreateToken from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCreateToken(key Key) *CreateToken {
	var createToken CreateToken
	if c.tryFindEntry(key, &createToken) {
		return &createToken
	}
	return nil
}
