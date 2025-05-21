package auth_models

import (
	"github.com/google/uuid"
)

type TokenMetadata struct {
	ClientID  uuid.UUID `json:"client_id"`
	AppID     int       `json:"app_id"`
	TokenType string    `json:"token_type"`
	Issuer    string    `json:"issuer"`
	Audiences []string  `json:"audiences"`
}
