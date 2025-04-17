package service

import (
	sl "backend/pkg/logger"
	"backend/service/apps/constants"
	"backend/service/apps/models"
	"context"
	"errors"
	"fmt"
	"log/slog"
)

func (s *Service) GetByCode(ctx context.Context, code string) (*models.App, error) {
	const op = "service.AppService.GetByCode"
	logger := s.logger.With(slog.String("op", op), slog.String("app_code", code))

	if err := validateCode(code, 50); err != nil {
		logger.Warn("Invalid code requested", slog.String("requested_code", code), sl.Err(err, false))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	app, err := s.provider.GetByCode(ctx, code)
	if err != nil {
		if errors.Is(err, constants.ErrNotFound) {
			logger.Warn("App not found")
			return nil, fmt.Errorf("%s: %w", op, constants.ErrNotFound)
		}
		logger.Error("Database error", slog.String("error", err.Error()), slog.String("error_type", "database"))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if app == nil {
		logger.Error("Unexpected nil app without error")
		return nil, fmt.Errorf("%s: internal server error", op)
	}

	logger.Info("App retrieved successfully", slog.Int("app_id", app.ID), slog.Bool("is_active", app.IsActive))
	return app, nil
}
