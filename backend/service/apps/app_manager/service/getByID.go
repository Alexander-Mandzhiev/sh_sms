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

func (s *Service) GetByID(ctx context.Context, id int) (*models.App, error) {
	const op = "service.AppService.GetByID"
	logger := s.logger.With(slog.String("op", op), slog.Int("app_id", id))

	if err := validateID(id); err != nil {
		logger.Warn("Invalid ID requested", slog.Int("requested_id", id), sl.Err(err, false))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	app, err := s.provider.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, handle.ErrNotFound) {
			logger.Warn("App not found")
			return nil, fmt.Errorf("%s: %w", op, handle.ErrNotFound)
		}
		logger.Error("Database error", slog.String("error", err.Error()), slog.String("error_type", "database"))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if app == nil {
		logger.Error("Unexpected nil app without error")
		return nil, fmt.Errorf("%s: internal server error", op)
	}

	logger.Info("App retrieved successfully", slog.String("app_code", app.Code), slog.Bool("is_active", app.IsActive))
	return app, nil
}
