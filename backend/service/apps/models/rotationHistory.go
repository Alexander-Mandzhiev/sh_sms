package models

import "time"

type RotationHistory struct {
	ClientID   string
	AppID      int
	SecretType string
	OldSecret  string
	NewSecret  string
	RotatedBy  string
	RotatedAt  time.Time
}
