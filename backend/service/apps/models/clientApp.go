package models

import "time"

type ClientApp struct {
	ClientID  string
	AppID     int
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
