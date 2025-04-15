package models

import "time"

type ClientApp struct {
	ClientID  string
	AppID     int32
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
