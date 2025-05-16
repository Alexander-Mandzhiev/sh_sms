package models

import (
	"backend/protos/gen/go/auth"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type TokenInfo struct {
	Active    bool
	ClientID  uuid.UUID
	UserID    uuid.UUID
	ExpiresAt time.Time
	IssuedAt  time.Time
	Scope     string
}

func TokenInfoToProto(ti *TokenInfo) *auth.TokenInfo {
	return &auth.TokenInfo{
		Active:   ti.Active,
		ClientId: ti.ClientID.String(),
		UserId:   ti.UserID.String(),
		Exp:      timestamppb.New(ti.ExpiresAt),
		Iat:      timestamppb.New(ti.IssuedAt),
		Scope:    ti.Scope,
	}
}

func TokenInfoFromProto(pb *auth.TokenInfo) (*TokenInfo, error) {
	clientID, err := uuid.Parse(pb.ClientId)
	if err != nil {
		return nil, fmt.Errorf("invalid client_id: %w", err)
	}

	userID, err := uuid.Parse(pb.UserId)
	if err != nil {
		return nil, fmt.Errorf("invalid user_id: %w", err)
	}

	return &TokenInfo{
		Active:    pb.Active,
		ClientID:  clientID,
		UserID:    userID,
		ExpiresAt: pb.Exp.AsTime(),
		IssuedAt:  pb.Iat.AsTime(),
		Scope:     pb.Scope,
	}, nil
}
