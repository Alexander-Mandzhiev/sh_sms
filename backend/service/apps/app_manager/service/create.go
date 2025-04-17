package service

import (
	"backend/service/apps/constants"
	"backend/service/apps/models"
	"context"
	"errors"
	"fmt"
	"log/slog"
)

func (s *Service) Create(ctx context.Context, req *models.CreateApp) (*models.App, error) {
	const op = "service.AppService.Create"
	logger := s.logger.With(slog.String("op", op), slog.String("code", req.Code), slog.String("name", req.Name))

	if err := validateCode(req.Code, 50); err != nil {
		logger.Warn("Validation failed", slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	if err := validateName(req.Name, 250); err != nil {
		logger.Warn("Validation failed", slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	app, err := s.provider.Create(ctx, req)
	if err != nil {
		if errors.Is(err, constants.ErrAlreadyExists) {
			logger.Warn("Code conflict", slog.String("code", req.Code))
			return nil, fmt.Errorf("%s: %w", op, constants.ErrAlreadyExists)
		}
		logger.Error("Create failed", slog.Any("error", err), slog.String("error_type", "database"))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	logger.Info("App created successfully", slog.Int("id", app.ID), slog.Bool("is_active", app.IsActive))
	return app, nil
}
