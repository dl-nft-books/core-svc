/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type NftPayment struct {
	Key
	Attributes    NftPaymentAttributes     `json:"attributes"`
	Relationships *NftPaymentRelationships `json:"relationships,omitempty"`
}
type NftPaymentResponse struct {
	Data     NftPayment `json:"data"`
	Included Included   `json:"included"`
}

type NftPaymentListResponse struct {
	Data     []NftPayment `json:"data"`
	Included Included     `json:"included"`
	Links    *Links       `json:"links"`
}

// MustNftPayment - returns NftPayment from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustNftPayment(key Key) *NftPayment {
	var nftPayment NftPayment
	if c.tryFindEntry(key, &nftPayment) {
		return &nftPayment
	}
	return nil
}
