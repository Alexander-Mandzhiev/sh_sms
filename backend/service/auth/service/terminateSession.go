package service

import (
	"backend/pkg/utils"
	"context"
	"fmt"
	"github.com/google/uuid"
	"log/slog"
)

func (s *AuthService) TerminateSession(ctx context.Context, sessionID uuid.UUID) error {
	const op = "service.TerminateSession"
	logger := s.logger.With(slog.String("op", op), slog.String("session_id", sessionID.String()))

	if err := utils.ValidateUUID(sessionID); err != nil {

	}
	logger.Debug("starting session termination")

	if err := s.session.RevokeSession(ctx, sessionID); err != nil {
		logger.Error("failed to revoke session", slog.Any("error", err))
		return fmt.Errorf("%s: %w", op, err)
	}

	logger.Debug("session terminated successfully")
	return nil
}
