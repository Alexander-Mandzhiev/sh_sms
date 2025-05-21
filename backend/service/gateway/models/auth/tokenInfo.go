package auth_models

import (
	"backend/pkg/utils"
	"backend/protos/gen/go/auth"
	"github.com/google/uuid"
	"time"
)

type TokenInfo struct {
	Active    bool      `json:"active"`
	ClientID  uuid.UUID `json:"client_id"`
	UserID    uuid.UUID `json:"user_id"`
	ExpiresAt time.Time `json:"expires_at"`
	IssuedAt  time.Time `json:"issued_at"`
	Scope     string    `json:"scope"`
}

func TokenInfoFromProto(pb *auth.TokenInfo) (*TokenInfo, error) {
	clientID, err := utils.ValidateAndReturnUUID(pb.GetClientId())
	if err != nil {
		return nil, err
	}

	userID, err := utils.ValidateAndReturnUUID(pb.GetUserId())
	if err != nil {
		return nil, err
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
