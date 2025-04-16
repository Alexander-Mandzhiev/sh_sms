package models

import "time"

type UpdateClientApp struct {
	ClientID  string
	AppID     int
	IsActive  *bool
	UpdatedAt time.Time
}
