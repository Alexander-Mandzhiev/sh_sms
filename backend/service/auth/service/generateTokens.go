package service

import (
	"backend/pkg/jwt_manager"
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

func (s *AuthService) generateTokens(ctx context.Context, userID uuid.UUID, clientID uuid.UUID,
	appID int, roles, permissions []string, tokenType jwt_manager.TokenType, duration time.Duration) (string, error) {
	const op = "auth.service.generateTokens"
	if userID == uuid.Nil {
		return "", fmt.Errorf("userID cannot be empty")
	}

	secret, err := s.getJWTSecret(ctx, clientID, appID, tokenType)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	claims := jwt_manager.CustomClaims{
		UserID:      userID,
		ClientID:    clientID,
		AppID:       appID,
		Roles:       roles,
		Permissions: permissions,
		TokenType:   tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		},
	}

	token, err := jwt_manager.Generate(claims, secret)
	if err != nil {
		return "", fmt.Errorf("token generation failed: %w", err)
	}

	return token, nil
}
