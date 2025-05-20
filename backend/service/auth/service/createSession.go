package service

import (
	"backend/pkg/jwt_manager"
	"backend/service/auth/models"
	"context"
	"fmt"
	"github.com/google/uuid"
	"net"
	"time"
)

func (s *AuthService) CreateSession(ctx context.Context, userId, clientId uuid.UUID, appId int, accessToken, refreshToken, ipAddress, userAgent string) error {

	const op = "auth.service.session.CreateSession"
	session := &models.Session{
		SessionID:        uuid.New(),
		UserID:           userId,
		ClientID:         clientId,
		AppID:            appId,
		AccessTokenHash:  jwt_manager.HashToken(accessToken),
		RefreshTokenHash: jwt_manager.HashToken(refreshToken),
		CreatedAt:        time.Now().UTC(),
		ExpiresAt:        time.Now().UTC().Add(s.cfg.JWT.AccessDuration),
		IPAddress:        net.ParseIP(ipAddress),
		UserAgent:        userAgent,
	}

	if err := s.session.CreateSession(ctx, session); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
