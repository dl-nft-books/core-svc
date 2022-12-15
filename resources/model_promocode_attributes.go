/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import (
	"time"
)

type PromocodeAttributes struct {
	// between 0.0 and 1.0 representing discount percentage
	Discount       float64         `json:"discount"`
	ExpirationDate time.Time       `json:"expiration_date"`
	Id             int64           `json:"id"`
	InitialUsages  int64           `json:"initial_usages"`
	LeftUsages     *int64          `json:"left_usages,omitempty"`
	Promocode      string          `json:"promocode"`
	State          *PromocodeState `json:"state,omitempty"`
}
