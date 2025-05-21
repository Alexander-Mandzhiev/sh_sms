package auth_models

import (
	"backend/protos/gen/go/auth"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type TokenValidationResult struct {
	Active    bool      `json:"active"`
	ClientID  uuid.UUID `json:"client_id"`
	UserID    uuid.UUID `json:"user_id"`
	ExpiresAt time.Time `json:"expires_at"`
	IssuedAt  time.Time `json:"issued_at"`
	Scope     string    `json:"scope"`
}

func TokenValidationToProto(res *TokenValidationResult) (*auth.TokenInfo, error) {
	if res.ClientID == uuid.Nil {
		return nil, fmt.Errorf("invalid client_id")
	}
	if res.UserID == uuid.Nil {
		return nil, fmt.Errorf("invalid user_id")
	}

	return &auth.TokenInfo{
		Active:   res.Active,
		ClientId: res.ClientID.String(),
		UserId:   res.UserID.String(),
		Exp:      timestamppb.New(res.ExpiresAt),
		Iat:      timestamppb.New(res.IssuedAt),
		Scope:    res.Scope,
	}, nil
}
