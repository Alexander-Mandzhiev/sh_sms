package auth_models

import (
	"backend/protos/gen/go/auth"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type TokenIntrospection struct {
	Active      bool               `json:"active"`
	ClientID    uuid.UUID          `json:"client_id"`
	UserID      uuid.UUID          `json:"user_id"`
	TokenType   string             `json:"token_type"`
	ExpiresAt   time.Time          `json:"expires_at"`
	IssuedAt    time.Time          `json:"issued_at"`
	Roles       []string           `json:"roles"`
	Permissions []string           `json:"permissions"`
	Metadata    IntrospectMetadata `json:"metadata"`
}

type IntrospectMetadata struct {
	IPAddress string `json:"ip_address"`
	UserAgent string `json:"user_agent"`
	ClientApp string `json:"client_app"`
	SessionID string `json:"session_id"`
	AppID     int    `json:"app_id"`
}

func TokenIntrospectionFromProto(pb *auth.TokenIntrospection) (*TokenIntrospection, error) {
	clientID, err := uuid.Parse(pb.ClientId)
	if err != nil {
		return nil, fmt.Errorf("invalid client_id: %w", err)
	}

	userID, err := uuid.Parse(pb.UserId)
	if err != nil {
		return nil, fmt.Errorf("invalid user_id: %w", err)
	}

	var metadata IntrospectMetadata
	if pb.Metadata != nil {
		metadata = IntrospectMetadata{
			IPAddress: pb.Metadata.IpAddress,
			UserAgent: pb.Metadata.UserAgent,
			ClientApp: pb.Metadata.ClientApp,
			SessionID: pb.Metadata.SessionId,
			AppID:     int(pb.Metadata.AppId),
		}
	}

	return &TokenIntrospection{
		Active:      pb.Active,
		ClientID:    clientID,
		UserID:      userID,
		TokenType:   pb.TokenType,
		ExpiresAt:   pb.Exp.AsTime(),
		IssuedAt:    pb.Iat.AsTime(),
		Roles:       pb.Roles,
		Permissions: pb.Permissions,
		Metadata:    metadata,
	}, nil
}
