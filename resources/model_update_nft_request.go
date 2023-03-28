/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type UpdateNftRequest struct {
	Key
	Attributes UpdateNftRequestAttributes `json:"attributes"`
}
type UpdateNftRequestRequest struct {
	Data     UpdateNftRequest `json:"data"`
	Included Included         `json:"included"`
}

type UpdateNftRequestListRequest struct {
	Data     []UpdateNftRequest `json:"data"`
	Included Included           `json:"included"`
	Links    *Links             `json:"links"`
}

// MustUpdateNftRequest - returns UpdateNftRequest from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustUpdateNftRequest(key Key) *UpdateNftRequest {
	var updateNftRequest UpdateNftRequest
	if c.tryFindEntry(key, &updateNftRequest) {
		return &updateNftRequest
	}
	return nil
}
