package service

import (
	sl "backend/pkg/logger"
	"backend/service/apps/models"
	"backend/service/utils"
	"context"
	"fmt"
	"log/slog"
)

func (s *Service) List(ctx context.Context, filter models.ListFilter) ([]models.App, int, error) {
	const op = "service.AppService.List"
	logger := s.logger.With(slog.String("op", op), slog.Int("page", filter.Page), slog.Int("count", filter.Count))
	if err := utils.ValidatePagination(filter.Page, filter.Count); err != nil {
		logger.Warn("Invalid pagination parameters", sl.Err(err, false))
		return nil, 0, fmt.Errorf("%s: %w", op, err)
	}
	if filter.IsActive != nil {
		logger = logger.With(slog.Bool("filter_active", *filter.IsActive))
	}
	apps, total, err := s.provider.List(ctx, filter)
	if err != nil {
		logger.Error("Failed to list applications", sl.Err(err, true), slog.Any("filter", filter))
		return nil, 0, fmt.Errorf("%s: %w", op, err)
	}
	if apps == nil {
		logger.Warn("Unexpected nil apps list")
		return make([]models.App, 0), total, nil
	}

	logger.Info("List operation completed", slog.Int("returned", len(apps)), slog.Int("total", total))
	return apps, total, nil
}
