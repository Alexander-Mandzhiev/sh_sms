package models

type CreateSecretParams struct {
	ClientID   string
	AppID      int
	SecretType string
	Algorithm  string
}
