package service

import (
	"backend/pkg/jwt_manager"
	"backend/pkg/utils"
	"backend/service/auth/handle"
	"backend/service/auth/models"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log/slog"
)

func (s *AuthService) findActiveSession(ctx context.Context, userID uuid.UUID, clientID uuid.UUID, appID int, refreshToken string) (*models.Session, error) {
	const op = "auth.service.session.findActiveSession"
	logger := s.logger.With(slog.String("op", op), slog.String("user_id", userID.String()), slog.String("client_id", clientID.String()), slog.Int("app_id", appID))

	logger.Debug("validating session parameters")

	if err := utils.ValidateAppID(appID); err != nil {
		logger.Warn("invalid app_id", slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if err := utils.ValidateUUID(clientID); err != nil {
		logger.Warn("invalid client_id", slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if err := utils.ValidateUUID(userID); err != nil {
		logger.Warn("invalid user_id", slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	session, err := s.session.GetSession(ctx, userID, clientID, appID, jwt_manager.HashToken(refreshToken))
	if err != nil {
		if errors.Is(err, handle.ErrSessionNotFound) {
			logger.Warn("session not found")
			return nil, handle.ErrInvalidToken
		}
		logger.Error("failed to get session", slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	logger.Debug("active session found", slog.String("session_id", session.SessionID.String()))
	return session, nil
}
