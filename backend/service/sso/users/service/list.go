package service

import (
	sl "backend/pkg/logger"
	"backend/pkg/utils"
	"backend/service/sso/models"
	"context"
	"fmt"
	"log/slog"
)

func (s *Service) List(ctx context.Context, req models.ListRequest) ([]models.User, int, error) {
	const op = "service.User.List"
	logger := s.logger.With(slog.String("op", op), slog.String("client_id", req.ClientID.String()), slog.Int("page", req.Page), slog.Int("count", req.Count))
	logger.Debug("starting users listing")

	if err := utils.ValidatePagination(req.Page, req.Count); err != nil {
		logger.Warn("Invalid pagination parameters", sl.Err(err, false))
		return nil, 0, fmt.Errorf("%s: %w", op, err)
	}

	users, total, err := s.provider.List(ctx, req)
	if err != nil {
		logger.Error("failed to list users", slog.Any("error", err))
		return nil, 0, fmt.Errorf("%w: database error", ErrInternal)
	}

	if len(users) == 0 {
		logger.Debug("no users found matching criteria")
		return []models.User{}, total, nil
	}

	logger.Debug("listing completed", slog.Int("total_users", total), slog.Int("returned_users", len(users)))
	return users, total, nil
}
