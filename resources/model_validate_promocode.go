/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type ValidatePromocode struct {
	Key
	Attributes    ValidatePromocodeAttributes    `json:"attributes"`
	Relationships ValidatePromocodeRelationships `json:"relationships"`
}
type ValidatePromocodeResponse struct {
	Data     ValidatePromocode `json:"data"`
	Included Included          `json:"included"`
}

type ValidatePromocodeListResponse struct {
	Data     []ValidatePromocode `json:"data"`
	Included Included            `json:"included"`
	Links    *Links              `json:"links"`
}

// MustValidatePromocode - returns ValidatePromocode from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustValidatePromocode(key Key) *ValidatePromocode {
	var validatePromocode ValidatePromocode
	if c.tryFindEntry(key, &validatePromocode) {
		return &validatePromocode
	}
	return nil
}
