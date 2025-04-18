package service

import (
	"backend/service/constants"
	"context"
	"errors"
	"fmt"
	"log/slog"
)

func (s *Service) DeleteByCode(ctx context.Context, code string) error {
	const op = "service.AppService.DeleteByCode"
	logger := s.logger.With(slog.String("op", op), slog.String("app_code", code))

	if err := validateCode(code, 50); err != nil {
		logger.Error("ID validation failed", slog.Any("error", err))
		return fmt.Errorf("%s: %w", op, err)
	}

	err := s.provider.DeleteByCode(ctx, code)
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
