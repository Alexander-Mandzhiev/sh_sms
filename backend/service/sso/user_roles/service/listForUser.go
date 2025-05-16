package service

import (
	"backend/pkg/utils"
	"backend/service/sso/models"
	"context"
	"fmt"
	"github.com/google/uuid"
	"log/slog"
)

func (s *Service) ListForUser(ctx context.Context, filter models.ListRequest) ([]*models.UserRole, int, error) {
	const op = "service.user_roles.ListForUser"
	logger := s.logger.With(slog.String("op", op), slog.String("user_id", filter.UserID.String()), slog.String("client_id", filter.ClientID.String()))
	logger.Debug("processing roles list request")

	if filter.UserID == nil || *filter.UserID == uuid.Nil {
		logger.Warn("empty user_id")
		return nil, 0, fmt.Errorf("%w: user_id", ErrInvalidArgument)
	}

	if filter.ClientID == nil || *filter.ClientID == uuid.Nil {
		logger.Warn("empty client_id")
		return nil, 0, fmt.Errorf("%w: client_id", ErrInvalidArgument)
	}

	if filter.AppID == nil || *filter.AppID < 0 {
		logger.Warn("empty app_id")
		return nil, 0, fmt.Errorf("%w: app_id", ErrInvalidArgument)
	}

	if err := utils.ValidateUUID(*filter.UserID); err != nil {
		logger.Warn("invalid user_id", slog.Any("error", err))
		return nil, 0, fmt.Errorf("%w: user_id", ErrInvalidArgument)
	}

	if err := utils.ValidateUUID(*filter.ClientID); err != nil {
		logger.Warn("invalid client_id", slog.Any("error", err))
		return nil, 0, fmt.Errorf("%w: client_id", ErrInvalidArgument)
	}

	if err := utils.ValidateAppID(*filter.AppID); err != nil {
		logger.Warn("invalid app_id", slog.Int("app_id", *filter.AppID))
		return nil, 0, fmt.Errorf("%w: app_id", ErrInvalidArgument)
	}

	if err := utils.ValidatePagination(filter.Page, filter.Count); err != nil {
		logger.Warn("invalid pagination", slog.Any("error", err), slog.Int("page", filter.Page), slog.Int("count", filter.Count))
		return nil, 0, fmt.Errorf("%w: %v", ErrInvalidArgument, err)
	}

	userExists, err := s.userProvider.Exists(ctx, *filter.ClientID, *filter.UserID)
	if err != nil {
		logger.Error("user existence check failed", slog.Any("error", err))
		return nil, 0, fmt.Errorf("%w: %v", ErrInternal, err)
	}
	if !userExists {
		logger.Warn("user not found")
		return nil, 0, ErrUserNotFound
	}

	roles, total, err := s.provider.ListForUser(ctx, filter)
	if err != nil {
		logger.Error("database operation failed", slog.Any("error", err))
		return nil, 0, fmt.Errorf("%w: %v", ErrInternal, err)
	}

	logger.Info("roles list retrieved", slog.Int("count", len(roles)), slog.Int("total", total))
	return roles, total, nil
}
