package models

type RotateSecretParams struct {
	ClientID   string
	AppID      int
	SecretType string
	RotatedBy  string
}
