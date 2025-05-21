package auth_models

import (
	"backend/pkg/utils"
	"backend/protos/gen/go/auth"
	"fmt"
	"time"
)

type AuthResponse struct {
	AccessToken string        `json:"access_token"`
	ExpiresIn   time.Time     `json:"expires_in"`
	User        UserInfo      `json:"user"`
	Metadata    TokenMetadata `json:"metadata"`
}

func AuthResponseFromProto(req *auth.AuthResponse) (*AuthResponse, error) {
	clientID, err := utils.ValidateAndReturnUUID(req.Metadata.ClientId)
	if err != nil {
		return nil, fmt.Errorf("failed to parse client_id: %w", err)
	}

	return &AuthResponse{
		AccessToken: req.AccessToken,
		ExpiresIn:   req.ExpiresAt.AsTime(),
		User: UserInfo{
			ID:          req.User.Id,
			Email:       req.User.Email,
			FullName:    req.User.FullName,
			IsActive:    req.User.IsActive,
			Roles:       req.User.Roles,
			Permissions: req.User.Permissions,
		},
		Metadata: TokenMetadata{
			ClientID:  clientID,
			AppID:     int(req.Metadata.AppId),
			TokenType: req.Metadata.TokenType,
			Issuer:    req.Metadata.Issuer,
			Audiences: req.Metadata.Audiences,
		},
	}, nil
}
