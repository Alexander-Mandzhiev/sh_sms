package service

import (
	sl "backend/pkg/logger"
	"backend/service/apps/app_manager/handle"
	"backend/service/apps/models"
	"context"
	"errors"
	"fmt"
	"log/slog"
)

func (s *Service) Update(ctx context.Context, id int, params models.UpdateApp) (*models.App, error) {
	const op = "service.AppService.Update"
	logger := s.logger.With(slog.String("op", op), slog.Int("app_id", id))

	if err := validateID(id); err != nil {
		logger.Warn("Invalid app ID", sl.Err(err, false))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if !params.HasUpdates() {
		logger.Warn("No fields to update")
		return nil, fmt.Errorf("%s: %w", op, handle.ErrNoUpdateFields)
	}

	if params.Code != nil {
		if err := validateCode(*params.Code, 50); err != nil {
			logger.Warn("Invalid code format", sl.Err(err, false))
			return nil, fmt.Errorf("%s: %w", op, err)
		}
	}

	if params.Name != nil {
		if err := validateName(*params.Name, 250); err != nil {
			logger.Warn("Invalid name format", sl.Err(err, false))
			return nil, fmt.Errorf("%s: %w", op, err)
		}
	}

	currentApp, err := s.provider.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, handle.ErrNotFound) {
			logger.Warn("App not found")
			return nil, fmt.Errorf("%s: %w", op, handle.ErrNotFound)
		}
		logger.Error("Failed to get current app", sl.Err(err, true))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	updatedApp := mergeAppUpdates(currentApp, params)
	result, err := s.provider.Update(ctx, updatedApp)
	if err != nil {
		if errors.Is(err, handle.ErrVersionConflict) {
			logger.Warn("Version conflict, retrying...")
			return s.Update(ctx, id, params)
		}
		logger.Error("Update failed", sl.Err(err, true))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	logger.Info("App updated successfully", slog.Int("new_version", result.Version))
	return result, nil

}

func mergeAppUpdates(current *models.App, updates models.UpdateApp) *models.App {
	updated := *current

	if updates.Code != nil {
		updated.Code = *updates.Code
	}
	if updates.Name != nil {
		updated.Name = *updates.Name
	}
	if updates.Description != nil {
		updated.Description = *updates.Description
	}
	if updates.IsActive != nil {
		updated.IsActive = *updates.IsActive
	}

	return &updated
}
