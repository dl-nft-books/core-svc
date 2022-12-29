/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type PriceAttributes struct {
	// amount of discount where 1% = 10^25
	Discount string `json:"discount"`
	// Timestamp when signature will expire
	EndTimestamp int64 `json:"end_timestamp"`
	// price per one token ($)
	Price string `json:"price"`
}
