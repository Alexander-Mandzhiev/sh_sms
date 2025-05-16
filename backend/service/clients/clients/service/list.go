package service

import (
	"backend/pkg/utils"
	"backend/service/clients/clients/handle"
	"backend/service/clients/clients/models"
	"context"
	"errors"
	"fmt"
	"log/slog"
	"strings"
)

func (s *Service) List(ctx context.Context, filter models.Filter, pagination models.Pagination) ([]*models.Client, int, error) {
	const op = "service.Clients.List"
	logger := s.logger.With(slog.String("op", op), slog.Int("page", pagination.Page), slog.Int("count", pagination.Count))
	logger.Debug("processing list request")

	if err := utils.ValidatePagination(pagination.Page, pagination.Count); err != nil {
		logger.Warn("invalid pagination", slog.String("error", err.Error()))
		return nil, 0, fmt.Errorf("%w: %s", handle.ErrInvalidArgument, err)
	}

	if filter.Search != nil {
		search := strings.TrimSpace(*filter.Search)
		if len(search) > 1000 {
			logger.Warn("search query too long")
			return nil, 0, handle.ErrInvalidArgument
		}
		if strings.Contains(*filter.Search, ";<>'\"\\") {
			logger.Warn("invalid characters in search query")
			return nil, 0, handle.ErrInvalidArgument
		}
		filter.Search = &search
	}

	if filter.TypeID != nil && *filter.TypeID <= 0 {
		logger.Warn("invalid type_id in filter")
		return nil, 0, fmt.Errorf("%w: type_id", handle.ErrInvalidArgument)
	}

	if filter.ActiveOnly == nil {
		defaultActive := true
		filter.ActiveOnly = &defaultActive
	}

	clients, total, err := s.provider.List(ctx, filter, pagination)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, 0, fmt.Errorf("%s: %w", op, handle.ErrTimeout)
		}
		logger.Error("database error", slog.String("error", err.Error()), slog.Any("filter", filter))
		return nil, 0, fmt.Errorf("%s: %w", op, handle.ErrInternal)
	}

	logger.Debug("list operation successful", slog.Int("total", total), slog.Int("returned", len(clients)))
	return clients, total, nil
}
