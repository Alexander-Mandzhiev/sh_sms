package service

import (
	"backend/service/apps/constants"
	"context"
	"errors"
	"fmt"
	"log/slog"
)

func (s *Service) DeleteByID(ctx context.Context, id int) error {
	const op = "service.AppService.DeleteByID"
	logger := s.logger.With(slog.String("op", op), slog.Int("app_id", id))

	if id <= 0 {
		logger.Warn("Invalid ID requested")
		return fmt.Errorf("%s: %w", op, constants.ErrInvalidID)
	}

	err := s.provider.DeleteByID(ctx, id)
	if err != nil {
		if errors.Is(err, constants.ErrNotFound) {
			logger.Warn("App not found for deletion")
			return fmt.Errorf("%s: %w", op, constants.ErrNotFound)
		}
		logger.Error("Deletion failed", slog.String("error", err.Error()), slog.String("error_type", "database"))
		return fmt.Errorf("%s: %w", op, err)
	}

	logger.Info("App deleted successfully")
	return nil
}
