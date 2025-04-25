package service

import (
	"backend/service/utils"
	"context"
	"fmt"
	"github.com/google/uuid"
	"log/slog"
)

func (s *Service) Delete(ctx context.Context, id uuid.UUID, appID int, permanent bool) error {
	const op = "service.Permission.Delete"
	logger := s.logger.With(slog.String("op", op), slog.String("permission_id", id.String()), slog.Int("app_id", appID), slog.Bool("permanent", permanent))
	logger.Debug("starting delete operation")

	if err := utils.ValidateUUID(id); err != nil {
		logger.Warn("invalid UUID format", slog.Any("error", err))
		return fmt.Errorf("%w: invalid permission id", ErrInvalidArgument)
	}

	if err := utils.ValidateAppID(appID); err != nil {
		logger.Warn("invalid app_id", slog.Any("error", err))
		return fmt.Errorf("%w: invalid app_id", ErrInvalidArgument)
	}

	exists, err := s.provider.ExistByID(ctx, id, appID)
	if err != nil {
		logger.Error("existence check failed", slog.Any("error", err))
		return fmt.Errorf("%s: %w", op, err)
	}
	if !exists {
		logger.Warn("permission not found")
		return fmt.Errorf("%w", ErrNotFound)
	}

	var deleteErr error
	if permanent {
		deleteErr = s.provider.HardDelete(ctx, id, appID)
	} else {
		deleteErr = s.provider.SoftDelete(ctx, id, appID)
	}

	if deleteErr != nil {
		logger.Error("delete operation failed", slog.Any("error", deleteErr), slog.Bool("permanent", permanent))
		return fmt.Errorf("%s: %w", op, deleteErr)
	}

	logger.Info("permission deleted successfully", slog.Bool("permanent", permanent), slog.String("id", id.String()))
	return nil
}
