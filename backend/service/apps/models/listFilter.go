package models

import "time"

type ListFilter struct {
	Page          int
	Count         int
	ClientID      *string
	AppID         *int
	IsActive      *bool
	SecretType    *string
	RotatedBy     *string
	RotatedAfter  *time.Time
	RotatedBefore *time.Time
}
