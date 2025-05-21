package service

import (
	"backend/pkg/utils"
	"backend/protos/gen/go/sso/users"
	"backend/service/auth/handle"
	"backend/service/auth/models"
	"context"
	"fmt"
	"log/slog"
)

func (s *AuthService) ListSessionsForUser(ctx context.Context, filter models.SessionFilter) ([]models.Session, error) {
	const op = "service.ListSessionsForUser"
	logger := s.logger.With(
		slog.String("op", op),
		slog.String("user_id", filter.UserID.String()),
		slog.String("client_id", filter.ClientID.String()),
		slog.Int("app_id", filter.AppID),
		slog.Int("page", filter.Page),
		slog.Int("count", filter.Count),
	)

	logger.Debug("starting sessions fetch for user")

	if err := utils.ValidateUUID(filter.UserID); err != nil {
		logger.Warn("invalid user_id", slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", op, handle.ErrInvalidUserID)
	}

	if err := utils.ValidateAppID(filter.AppID); err != nil {
		logger.Warn("invalid app_id", slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", op, handle.ErrInvalidAppID)
	}

	if err := utils.ValidatePagination(filter.Page, filter.Count); err != nil {
		logger.Warn("invalid pagination", slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", op, handle.ErrInvalidPagination)
	}

	user, err := s.users.GetUser(ctx, &users.GetRequest{Id: filter.UserID.String(), ClientId: filter.ClientID.String()})
	if err != nil {
		logger.Warn("failed to get user", slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	sessions, err := s.session.ListSessionsForUser(ctx, filter, user.FullName, user.Phone, user.Email)
	if err != nil {
		logger.Error("failed to list user sessions", slog.Any("error", err), slog.Int("attempt", 1))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	logger.Info("successfully fetched user sessions", slog.Int("session_count", len(sessions)))
	return sessions, nil
}
