package service

import (
	"backend/service/constants"
	"backend/service/sso/models"
	"backend/service/utils"
	"context"
	"fmt"
	"log/slog"
)

func (s *Service) List(ctx context.Context, filter models.ListRequest) ([]models.Permission, int, error) {
	const op = "service.Permission.List"
	logger := s.logger.With(slog.String("op", op), slog.Any("filter", filter))
	logger.Debug("processing list request")

	if filter.AppID == nil || *filter.AppID <= 0 {
		logger.Warn("invalid app_id")
		return nil, 0, fmt.Errorf("%w: app_id is required", ErrInvalidArgument)
	}

	if err := utils.ValidatePagination(filter.Page, filter.Count); err != nil {
		logger.Warn("invalid pagination", slog.Any("error", err), slog.Int("page", filter.Page), slog.Int("count", filter.Count))
		return nil, 0, fmt.Errorf("%w: %v", constants.ErrInvalidArgument, err)
	}

	if filter.CodeFilter != nil && len(*filter.CodeFilter) > 100 {
		logger.Warn("code filter too long")
		return nil, 0, fmt.Errorf("%w: code filter exceeds limit", ErrInvalidArgument)
	}
	if filter.CategoryFilter != nil && len(*filter.CategoryFilter) > 50 {
		logger.Warn("category filter too long")
		return nil, 0, fmt.Errorf("%w: code filter exceeds limit", ErrInvalidArgument)
	}

	permissions, total, err := s.provider.List(ctx, filter)
	if err != nil {
		logger.Error("database operation failed", slog.Any("error", err), slog.Int("app_id", *filter.AppID))
		return nil, 0, fmt.Errorf("%s: %w", op, err)
	}

	logger.Info("list operation completed", slog.Int("result_count", len(permissions)), slog.Int("total_records", total))
	return permissions, total, nil
}
