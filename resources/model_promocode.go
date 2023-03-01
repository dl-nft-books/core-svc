/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type Promocode struct {
	Key
	Attributes PromocodeAttributes `json:"attributes"`
}
type PromocodeResponse struct {
	Data     Promocode `json:"data"`
	Included Included  `json:"included"`
}

type PromocodeListResponse struct {
	Data     []Promocode `json:"data"`
	Included Included    `json:"included"`
	Links    *Links      `json:"links"`
}

// MustPromocode - returns Promocode from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustPromocode(key Key) *Promocode {
	var promocode Promocode
	if c.tryFindEntry(key, &promocode) {
		return &promocode
	}
	return nil
}
