/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type UpdateToken struct {
	Key
	Attributes UpdateTokenAttributes `json:"attributes"`
}
type UpdateTokenRequest struct {
	Data     UpdateToken `json:"data"`
	Included Included    `json:"included"`
}

type UpdateTokenListRequest struct {
	Data     []UpdateToken `json:"data"`
	Included Included      `json:"included"`
	Links    *Links        `json:"links"`
}

// MustUpdateToken - returns UpdateToken from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustUpdateToken(key Key) *UpdateToken {
	var updateToken UpdateToken
	if c.tryFindEntry(key, &updateToken) {
		return &updateToken
	}
	return nil
}
