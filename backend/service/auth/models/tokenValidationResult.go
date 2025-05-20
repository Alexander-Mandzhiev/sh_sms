package models

import (
	"backend/protos/gen/go/auth"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type TokenValidationResult struct {
	Active    bool
	ClientID  uuid.UUID
	UserID    uuid.UUID
	ExpiresAt time.Time
	IssuedAt  time.Time
	Scope     string
}

func TokenValidationToProto(res *TokenValidationResult) *auth.TokenInfo {
	return &auth.TokenInfo{
		Active:   res.Active,
		ClientId: res.ClientID.String(),
		UserId:   res.UserID.String(),
		Exp:      timestamppb.New(res.ExpiresAt),
		Iat:      timestamppb.New(res.IssuedAt),
		Scope:    res.Scope,
	}
}
