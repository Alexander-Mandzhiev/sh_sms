package service

import (
	"backend/service/clients/client_types/models"
	"context"
	"errors"
	"fmt"
	"log/slog"
)

func (s *Service) Restore(ctx context.Context, id int) (*models.ClientType, error) {
	const op = "service.ClientType.Restore"
	logger := s.logger.With(slog.String("op", op), slog.Int("id", id))
	logger.Debug("starting restore operation")

	if id <= 0 {
		err := fmt.Errorf("%w: invalid ID", ErrInvalidArgument)
		logger.Warn("validation failed")
		return nil, err
	}

	restoredType, err := s.provider.Restore(ctx, id)
	if err != nil {
		logger.Error("restore operation failed", slog.Any("error", err), slog.String("stage", "repository_call"))
		if errors.Is(err, ErrNotFound) {
			return nil, fmt.Errorf("%w: %d", ErrNotFound, id)
		}
		if errors.Is(err, ErrAlreadyActive) {
			return nil, fmt.Errorf("%w: %d", ErrAlreadyActive, id)
		}
		if errors.Is(err, ErrCodeConflict) {
			return nil, fmt.Errorf("%w: code conflict", ErrConflict)
		}
		return nil, fmt.Errorf("%w: restore failed", ErrInternal)
	}

	logger.Info("client type restored successfully", slog.Bool("is_active", restoredType.IsActive), slog.String("code", restoredType.Code))
	return restoredType, nil
}
