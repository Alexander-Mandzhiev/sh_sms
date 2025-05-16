package service

import (
	"backend/pkg/jwt_manager"
	"backend/service/auth/models"
	"context"
	"github.com/google/uuid"
	"time"
)

func (s *AuthService) createSession(ctx context.Context, userId, clientId uuid.UUID, appId int, accessToken, refreshToken string) error {
	_ = &models.Session{
		SessionID:        uuid.New(),
		UserID:           userId,
		ClientID:         clientId,
		AppID:            appId,
		AccessTokenHash:  jwt_manager.HashToken(accessToken),
		RefreshTokenHash: jwt_manager.HashToken(refreshToken),
		CreatedAt:        time.Now(),
		ExpiresAt:        time.Now().Add(24 * time.Hour),
		//IPAddress:        req.IPAddress,
		//UserAgent:        req.UserAgent,
	}

	return nil
}
