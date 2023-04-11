/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "time"

type CreatePromocodeAttributes struct {
	// Book with can support this promocode
	Books *[]int64 `json:"books,omitempty"`
	// between 0.0 and 1.0 representing discount percentage
	Discount float64 `json:"discount"`
	// Time of expiration
	ExpirationDate time.Time `json:"expiration_date"`
	// how many times you can use promocode
	InitialUsages int64 `json:"initial_usages"`
	// Custom promocode
	Promocode *string `json:"promocode,omitempty"`
}
