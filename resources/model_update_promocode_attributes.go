/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "time"

type UpdatePromocodeAttributes struct {
	// Book with can support this promocode
	Books *[]int64 `json:"books,omitempty"`
	// between 0.0 and 1.0 representing discount percentage
	Discount *float64 `json:"discount,omitempty"`
	// Time of expiration
	ExpirationDate *time.Time `json:"expiration_date,omitempty"`
	// how many times you can use promocode
	InitialUsages *int64 `json:"initial_usages,omitempty"`
	// Custom promocode
	Promocode *string `json:"promocode,omitempty"`
	// promocode status
	State *PromocodeState `json:"state,omitempty"`
	// how many times promocode has been used
	Usages *int64 `json:"usages,omitempty"`
}
