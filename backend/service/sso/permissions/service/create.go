package service

import (
	"backend/pkg/utils"
	"backend/service/sso/models"
	"context"
	"fmt"
	"log/slog"
)

func (s *Service) Create(ctx context.Context, perm models.Permission) (*models.Permission, error) {
	const op = "service.Permissions.Create"
	logger := s.logger.With(slog.String("op", op), slog.Int("app_id", perm.AppID), slog.String("code", perm.Code), slog.String("category", perm.Category))
	logger.Debug("validating input")

	if err := utils.ValidateString(perm.Code, 100); err != nil {
		logger.Warn("invalid app ID", slog.Any("error", err))
		return nil, fmt.Errorf("%w: code is required", ErrInvalidArgument)
	}

	if err := utils.ValidateAppID(perm.AppID); err != nil {
		logger.Warn("invalid app ID", slog.Any("error", err))
		return nil, fmt.Errorf("%w: invalid app_id", ErrInvalidArgument)
	}

	exists, err := s.provider.ExistByCode(ctx, perm.Code, perm.AppID)
	if err != nil {
		logger.Error("code check failed", slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	if exists {
		logger.Warn("duplicate code")
		return nil, fmt.Errorf("%w: code already exists", ErrAlreadyExists)
	}

	perm.IsActive = true
	perm.DeletedAt = nil

	logger.Debug("creating permission in repository")
	createdPerm, err := s.provider.Create(ctx, perm)
	if err != nil {
		logger.Error("repository create failed", slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	logger.Info("permission created successfully", slog.String("new_id", createdPerm.ID.String()), slog.Time("created_at", createdPerm.CreatedAt))
	return createdPerm, nil
}
