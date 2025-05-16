package models

import (
	"backend/protos/gen/go/auth"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type TokenIntrospection struct {
	Active         bool
	ClientID       uuid.UUID
	UserID         uuid.UUID
	TokenType      string
	ExpiresAt      time.Time
	IssuedAt       time.Time
	Roles          []string
	Permissions    []string
	ClientMetadata map[string]string
}

func TokenIntrospectionToProto(ti *TokenIntrospection) *auth.TokenIntrospection {
	return &auth.TokenIntrospection{
		Active:      ti.Active,
		ClientId:    ti.ClientID.String(),
		UserId:      ti.UserID.String(),
		TokenType:   ti.TokenType,
		Exp:         timestamppb.New(ti.ExpiresAt),
		Iat:         timestamppb.New(ti.IssuedAt),
		Roles:       ti.Roles,
		Permissions: ti.Permissions,
		Metadata:    ti.ClientMetadata,
	}
}
