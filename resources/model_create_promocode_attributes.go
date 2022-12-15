/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import (
	"time"
)

type CreatePromocodeAttributes struct {
	// between 0.0 and 1.0 representing discount percentage
	Discount       float64   `json:"discount"`
	ExpirationDate time.Time `json:"expiration_date"`
	InitialUsages  int64     `json:"initial_usages"`
}
