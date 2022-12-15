/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type CreatePromocode struct {
	Key
	Attributes CreatePromocodeAttributes `json:"attributes"`
}
type CreatePromocodeRequest struct {
	Data     CreatePromocode `json:"data"`
	Included Included        `json:"included"`
}

type CreatePromocodeListRequest struct {
	Data     []CreatePromocode `json:"data"`
	Included Included          `json:"included"`
	Links    *Links            `json:"links"`
}

// MustCreatePromocode - returns CreatePromocode from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCreatePromocode(key Key) *CreatePromocode {
	var createPromocode CreatePromocode
	if c.tryFindEntry(key, &createPromocode) {
		return &createPromocode
	}
	return nil
}
