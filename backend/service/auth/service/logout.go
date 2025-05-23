package service

import (
	"backend/pkg/jwt_manager"
	"backend/service/auth/handle"
	"context"
	"fmt"
	"log/slog"
)

func (s *AuthService) Logout(ctx context.Context, refreshToken string) error {
	const op = "service.Logout"
	logger := s.logger.With(slog.String("op", op), slog.String("refresh_token_hash", jwt_manager.HashToken(refreshToken)))

	if refreshToken == "" {
		logger.Warn("empty tokens provided")
		return fmt.Errorf("%s: %w", op, handle.ErrInvalidToken)
	}

	session, err := s.session.GetSessionByToken(ctx, jwt_manager.HashToken(refreshToken))
	if err != nil {
		logger.Error("session lookup failed", slog.Any("error", err))
		return fmt.Errorf("%s: %w", op, err)
	}

	logger = logger.With(slog.String("session_id", session.SessionID.String()), slog.String("user_id", session.UserID.String()), slog.String("client_id", session.ClientID.String()), slog.Int("app_id", session.AppID))

	if err = s.session.RevokeSession(ctx, session.SessionID); err != nil {
		logger.Error("session revocation failed", slog.Any("error", err))
		return fmt.Errorf("%s: %w", op, err)
	}

	logger.Info("session revoked successfully")
	return nil
}
