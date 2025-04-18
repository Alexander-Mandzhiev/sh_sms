package models

import "time"

type RotationHistory struct {
	ID         int       `db:"id"`
	ClientID   string    `db:"client_id"`
	AppID      int       `db:"app_id"`
	SecretType string    `db:"secret_type"`
	OldSecret  string    `db:"old_secret"`
	NewSecret  string    `db:"new_secret"`
	RotatedBy  string    `db:"rotated_by"`
	RotatedAt  time.Time `db:"rotated_at"`
}
