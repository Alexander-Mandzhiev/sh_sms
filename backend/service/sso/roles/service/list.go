package service

import (
	"backend/service/sso/models"
	"backend/service/utils"
	"context"
	"fmt"
	"github.com/google/uuid"
	"log/slog"
)

func (s *Service) List(ctx context.Context, req models.ListRequest) ([]models.Role, int, error) {
	const op = "service.Roles.List"
	logger := s.logger.With(slog.String("op", op), slog.String("client_id", req.ClientID.String()), slog.Int("page", req.Page), slog.Int("count", req.Count))
	logger.Debug("attempting to list roles")

	if req.ClientID == nil || *req.ClientID == uuid.Nil {
		logger.Warn("empty client_id")
		return nil, 0, fmt.Errorf("%w: client_id", ErrInvalidArgument)
	}

	if err := utils.ValidatePagination(req.Page, req.Count); err != nil {
		logger.Warn("invalid pagination", slog.Any("error", err), slog.Int("page", req.Page), slog.Int("count", req.Count))
		return nil, 0, fmt.Errorf("%w: %v", ErrInvalidArgument, err)
	}

	if req.NameFilter != nil && *req.NameFilter != "" {
		if err := utils.ValidateRoleName(*req.NameFilter, 150); err != nil {
			logger.Warn("invalid name filter", slog.String("filter", *req.NameFilter), slog.Any("error", err))
			return nil, 0, fmt.Errorf("%w: name_filter", ErrInvalidArgument)
		}
	}

	if req.LevelFilter != nil && *req.LevelFilter < 0 {
		logger.Warn("invalid level filter", slog.Int("level", *req.LevelFilter))
		return nil, 0, fmt.Errorf("%w: level_filter", ErrInvalidArgument)
	}

	rolesList, total, err := s.provider.List(ctx, req)
	if err != nil {
		logger.Error("database error", slog.Any("error", err))
		return nil, 0, fmt.Errorf("%w: %v", ErrInternal, err)
	}

	logger.Debug("successfully listed roles", slog.Int("count", len(rolesList)), slog.Int("total", total))
	return rolesList, total, nil
}
