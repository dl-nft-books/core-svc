/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type UpdatePromocode struct {
	Key
	Attributes UpdatePromocodeAttributes `json:"attributes"`
}
type UpdatePromocodeRequest struct {
	Data     UpdatePromocode `json:"data"`
	Included Included        `json:"included"`
}

type UpdatePromocodeListRequest struct {
	Data     []UpdatePromocode `json:"data"`
	Included Included          `json:"included"`
	Links    *Links            `json:"links"`
}

// MustUpdatePromocode - returns UpdatePromocode from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustUpdatePromocode(key Key) *UpdatePromocode {
	var updatePromocode UpdatePromocode
	if c.tryFindEntry(key, &updatePromocode) {
		return &updatePromocode
	}
	return nil
}
