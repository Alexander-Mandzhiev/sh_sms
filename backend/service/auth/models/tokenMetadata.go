package models

import (
	"backend/protos/gen/go/auth"
	"github.com/google/uuid"
)

type TokenMetadata struct {
	ClientID  uuid.UUID
	AppID     int
	TokenType string
	Issuer    string
	Audiences []string
}

func TokenMetadataToProto(tm *TokenMetadata) *auth.TokenMetadata {
	return &auth.TokenMetadata{
		ClientId:  tm.ClientID.String(),
		AppId:     int32(tm.AppID),
		TokenType: tm.TokenType,
		Issuer:    tm.Issuer,
		Audiences: tm.Audiences,
	}
}
