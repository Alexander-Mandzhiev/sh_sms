package service

import (
	sl "backend/pkg/logger"
	"backend/service/apps/client_apps/handle"
	"backend/service/apps/models"
	"context"
	"fmt"
	"log/slog"
)

func (s *Service) List(ctx context.Context, filter models.ListFilter) ([]*models.ClientApp, int, error) {
	const op = "service.ClientApp.List"
	logger := s.logger.With(slog.String("op", op), slog.Int("page", filter.Page), slog.Int("count", filter.Count))

	if err := validatePagination(filter.Page, filter.Count); err != nil {
		logger.Warn("invalid pagination parameters", sl.Err(err, false))
		return nil, 0, fmt.Errorf("%w: %v", handle.ErrInvalidArgument, err)
	}

	if filter.ClientID != nil {
		if err := validateClientID(*filter.ClientID); err != nil {
			logger.Warn("invalid client_id filter", slog.String("client_id", *filter.ClientID), sl.Err(err, false))
			return nil, 0, fmt.Errorf("%w: %v", handle.ErrInvalidArgument, err)
		}
	}

	if filter.AppID != nil {
		if err := validateAppID(*filter.AppID); err != nil {
			logger.Warn("invalid app_id filter", slog.Int("app_id", *filter.AppID), sl.Err(err, false))
			return nil, 0, fmt.Errorf("%w: %v", handle.ErrInvalidArgument, err)
		}
	}

	apps, total, err := s.provider.List(ctx, filter)
	if err != nil {
		logger.Error("failed to list client apps", sl.Err(err, false))
		return nil, 0, s.convertError(err)
	}

	logger.Info("successfully listed client apps", slog.Int("returned", len(apps)), slog.Int("total", total))

	return apps, total, nil
}
