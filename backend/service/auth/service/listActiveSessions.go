package service

import (
	"backend/pkg/utils"
	"backend/service/auth/models"
	"context"
	"fmt"
	"log/slog"
)

func (s *AuthService) ListActiveSessions(ctx context.Context, filter models.SessionFilter) ([]models.Session, error) {
	const op = "service.ListActiveSessions"
	logger := s.logger.With(slog.String("op", op), slog.String("user_id", filter.UserID.String()), slog.String("client_id", filter.ClientID.String()), slog.Int("app_id", filter.AppID), slog.Int("page", filter.Page), slog.Int("count", filter.Count))

	logger.Debug("fetching active sessions")

	if err := utils.ValidateAppID(filter.AppID); err != nil {
		logger.Warn("invalid app_id", slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if err := utils.ValidateUUID(filter.ClientID); err != nil {
		logger.Warn("invalid client_id", slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if err := utils.ValidateUUID(filter.UserID); err != nil {
		logger.Warn("invalid user_id", slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if err := utils.ValidatePagination(filter.Page, filter.Count); err != nil {
		logger.Warn("invalid page", slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	sessions, err := s.session.ListSessions(ctx, filter)
	if err != nil {
		logger.Error("failed to list sessions", slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	logger.Debug("sessions fetched successfully", slog.Int("count", len(sessions)))
	return sessions, nil
}
