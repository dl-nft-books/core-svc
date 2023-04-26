/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "time"

type PromocodeAttributes struct {
	// Book with can support this promocode
	Books []int64 `json:"books"`
	// between 0.0 and 1.0 representing discount percentage
	Discount float64 `json:"discount"`
	// Time of expiration
	ExpirationDate time.Time `json:"expiration_date"`
	Id             int64     `json:"id"`
	// how many times you can use promocode
	InitialUsages int64  `json:"initial_usages"`
	Promocode     string `json:"promocode"`
	// promocode status
	State PromocodeState `json:"state"`
	// how many times promocode has been used
	Usages int64 `json:"usages"`
}
