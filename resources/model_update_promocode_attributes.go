/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import (
	"time"
)

type UpdatePromocodeAttributes struct {
	// between 0.0 and 1.0 representing discount percentage
	Discount       *float64        `json:"discount,omitempty"`
	ExpirationDate *time.Time      `json:"expiration_date,omitempty"`
	InitialUsages  *int64          `json:"initial_usages,omitempty"`
	State          *PromocodeState `json:"state,omitempty"`
	Usages         *int64          `json:"usages,omitempty"`
}
