package service

import (
	"backend/service/sso/models"
	"backend/service/utils"
	"context"
	"fmt"
	"github.com/google/uuid"
	"log/slog"
)

func (s *Service) ListForRole(ctx context.Context, filter models.ListRequest) ([]*models.UserRole, int, error) {
	const op = "service.user_roles.ListForRole"
	logger := s.logger.With(slog.String("op", op), slog.String("role_id", filter.RoleID.String()), slog.String("client_id", filter.ClientID.String()))
	logger.Debug("processing users list for role")

	if filter.ClientID == nil || *filter.ClientID == uuid.Nil {
		logger.Warn("empty client_id")
		return nil, 0, fmt.Errorf("%w: client_id", ErrInvalidArgument)
	}
	if filter.RoleID == nil || *filter.RoleID == uuid.Nil {
		logger.Warn("empty role_id")
		return nil, 0, fmt.Errorf("%w: role_id", ErrInvalidArgument)
	}
	if filter.AppID == nil || *filter.AppID < 0 {
		logger.Warn("empty app_id")
		return nil, 0, fmt.Errorf("%w: app_id", ErrInvalidArgument)
	}

	if err := utils.ValidateUUID(*filter.RoleID); err != nil {
		logger.Warn("invalid role_id", slog.Any("error", err))
		return nil, 0, fmt.Errorf("%w: role_id", ErrInvalidArgument)
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

	roleExists, err := s.roleProvider.RoleExistsByID(ctx, *filter.ClientID, *filter.RoleID, *filter.AppID)
	if err != nil {
		logger.Error("role existence check failed", slog.Any("error", err))
		return nil, 0, fmt.Errorf("%w: %v", ErrInternal, err)
	}
	if !roleExists {
		logger.Warn("role not found")
		return nil, 0, ErrRoleNotFound
	}

	users, total, err := s.provider.ListForRole(ctx, filter)
	if err != nil {
		logger.Error("database operation failed", slog.Any("error", err), slog.Any("filter", filter))
		return nil, 0, fmt.Errorf("%w: %v", ErrInternal, err)
	}

	logger.Info("users list retrieved", slog.Int("count", len(users)), slog.Int("total", total))
	return users, total, nil
}
