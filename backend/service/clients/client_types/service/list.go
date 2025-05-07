package service

import (
	"backend/service/clients/client_types/models"
	"backend/service/utils"
	"context"
	"fmt"
	"log/slog"
	"strings"
)

func (s *Service) List(ctx context.Context, filter models.Filter, pagination models.Pagination) ([]*models.ClientType, int, error) {
	const op = "service.ClientType.List"
	logger := s.logger.With(slog.String("op", op), slog.Any("filter", filter), slog.Any("pagination", pagination))
	logger.Debug("processing list request")

	if err := utils.ValidatePagination(pagination.Page, pagination.PageSize); err != nil {
		logger.Warn("invalid pagination parameters", slog.Any("error", err))
		return nil, 0, fmt.Errorf("%w: %v", ErrInvalidArgument, err)
	}

	if filter.Search != nil {
		if len(*filter.Search) > 1000 {
			logger.Warn("search parameter exceeds 100 characters")
			return nil, 0, fmt.Errorf("search query too long")
		}
		if strings.Contains(*filter.Search, ";<>'\"\\") {
			logger.Warn("search parameter contains semicolon")
			return nil, 0, fmt.Errorf("invalid characters in search")
		}
	}

	activeOnly := true
	if filter.ActiveOnly != nil {
		activeOnly = *filter.ActiveOnly
	}
	filter.ActiveOnly = &activeOnly

	clientTypes, total, err := s.provider.List(ctx, filter, pagination)
	if err != nil {
		logger.Error("failed to list client types", slog.Any("error", err), slog.String("stage", "repository_call"))
		return nil, 0, fmt.Errorf("%w: %v", ErrInternal, err)
	}

	logger.Info("successfully retrieved client types", slog.Int("returned", len(clientTypes)), slog.Int("total", total))
	return clientTypes, total, nil
}
