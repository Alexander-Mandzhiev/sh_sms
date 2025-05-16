package service

import (
	"backend/pkg/utils"
	"context"
	"fmt"
	"log/slog"

	"backend/service/sso/models"
)

func (s *Service) Update(ctx context.Context, updatedPerm models.Permission) (*models.Permission, error) {
	const op = "service.Permission.Update"
	logger := s.logger.With(slog.String("op", op), slog.String("id", updatedPerm.ID.String()), slog.Int("app_id", updatedPerm.AppID))
	logger.Debug("starting update operation")

	if err := utils.ValidateUUID(updatedPerm.ID); err != nil {
		logger.Warn("invalid UUID format")
		return nil, fmt.Errorf("%w: invalid permission id", ErrInvalidArgument)
	}

	if updatedPerm.AppID <= 0 {
		logger.Warn("invalid app_id")
		return nil, fmt.Errorf("%w: invalid app_id", ErrInvalidArgument)
	}

	current, err := s.provider.GetByID(ctx, updatedPerm.ID, updatedPerm.AppID)
	if err != nil {
		logger.Error("failed to get current version", slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var changes bool
	logger = logger.With(slog.Any("current", current), slog.Any("update", updatedPerm))

	if updatedPerm.Code != "" && updatedPerm.Code != current.Code {
		if len(updatedPerm.Code) > 100 {
			logger.Warn("code too long")
			return nil, fmt.Errorf("%w: code exceeds 100 characters", ErrInvalidArgument)
		}
		var exists bool
		exists, err = s.provider.ExistByCode(ctx, updatedPerm.Code, updatedPerm.AppID)
		if err != nil {
			logger.Error("code check failed", slog.Any("error", err))
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		if exists {
			logger.Warn("duplicate code detected")
			return nil, fmt.Errorf("%w: code already exists", ErrAlreadyExists)
		}
		changes = true
		logger.Debug("code changed", slog.String("old", current.Code), slog.String("new", updatedPerm.Code))
	}

	if updatedPerm.Category != "" && updatedPerm.Category != current.Category {
		if len(updatedPerm.Category) > 50 {
			logger.Warn("category too long")
			return nil, fmt.Errorf("%w: category exceeds 50 characters", ErrInvalidArgument)
		}
		changes = true
		logger.Debug("category changed", slog.String("old", current.Category), slog.String("new", updatedPerm.Category))
	}

	if updatedPerm.Description != current.Description {
		changes = true
		logger.Debug("description changed", slog.String("old", current.Description), slog.String("new", updatedPerm.Description))
	}

	if !changes {
		logger.Info("no changes detected")
		return current, nil
	}

	result, err := s.provider.Update(ctx, updatedPerm)
	if err != nil {
		logger.Error("update failed", slog.Any("error", err), slog.Any("update_data", updatedPerm))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	logger.Info("permission updated successfully", slog.Time("updated_at", result.UpdatedAt), slog.String("new_code", result.Code))
	return result, nil
}
