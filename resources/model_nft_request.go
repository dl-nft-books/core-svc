/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type NftRequest struct {
	Key
	Attributes    NftRequestAttributes    `json:"attributes"`
	Relationships NftRequestRelationships `json:"relationships"`
}
type NftRequestResponse struct {
	Data     NftRequest `json:"data"`
	Included Included   `json:"included"`
}

type NftRequestListResponse struct {
	Data     []NftRequest `json:"data"`
	Included Included     `json:"included"`
	Links    *Links       `json:"links"`
}

// MustNftRequest - returns NftRequest from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustNftRequest(key Key) *NftRequest {
	var nftRequest NftRequest
	if c.tryFindEntry(key, &nftRequest) {
		return &nftRequest
	}
	return nil
}
