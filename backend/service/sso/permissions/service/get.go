package service

import (
	"backend/pkg/utils"
	"backend/service/sso/models"
	"backend/service/sso/permissions/repository"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log/slog"
)

func (s *Service) Get(ctx context.Context, id uuid.UUID, appID int) (*models.Permission, error) {
	const op = "service.Permission.Get"
	logger := s.logger.With(slog.String("op", op), slog.String("permission_id", id.String()), slog.Int("app_id", appID))
	logger.Debug("attempting to get permission")

	if err := utils.ValidateUUID(id); err != nil {
		logger.Warn("invalid UUID format", slog.Any("error", err))
		return nil, fmt.Errorf("%w: invalid permission id", ErrInvalidArgument)
	}

	if err := utils.ValidateAppID(appID); err != nil {
		logger.Warn("app_id validation failed", slog.Any("error", err))
		return nil, fmt.Errorf("%w: app_id must be positive", ErrInvalidArgument)
	}

	perm, err := s.provider.GetByID(ctx, id, appID)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			logger.Warn("permission not found", slog.String("id", id.String()))
			return nil, fmt.Errorf("%w: %s", ErrNotFound, err)
		}
		logger.Error("database error", slog.Any("error", err), slog.String("id", id.String()))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	logger.Info("permission retrieved successfully", slog.String("code", perm.Code), slog.Bool("is_active", perm.IsActive))
	return perm, nil
}
