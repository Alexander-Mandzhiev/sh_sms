package models

import "time"

type Secret struct {
	ClientID      string
	AppID         int
	SecretType    string // "access" или "refresh"
	CurrentSecret string
	Algorithm     string
	SecretVersion int
	GeneratedAt   time.Time
	RevokedAt     *time.Time // nil если не отозван
}
