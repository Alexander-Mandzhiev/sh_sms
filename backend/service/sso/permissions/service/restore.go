package service

import (
	"backend/service/sso/models"
	"context"
	"fmt"
	"github.com/google/uuid"
	"log/slog"
	"time"
)

func (s *Service) Restore(ctx context.Context, id uuid.UUID, appID int) (*models.Permission, error) {
	const op = "service.Permission.Restore"
	logger := s.logger.With(slog.String("op", op), slog.Any("permission_id", id), slog.Int("app_id", appID))
	logger.Debug("attempting to restore permission")

	existingPerm, err := s.provider.ExistByID(ctx, id, appID)
	if err != nil {
		logger.Error("failed to check permission existence", slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if !existingPerm {
		logger.Warn("attempted to restore non-deleted permission")
		return nil, fmt.Errorf("%w: permission not deleted", ErrInvalidState)
	}

	if err = s.provider.Restore(ctx, id, appID); err != nil {
		logger.Error("provider restore failed", slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	restoredPerm, err := s.provider.GetByID(ctx, id, appID)
	if err != nil {
		logger.Error("failed to fetch restored permission", slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	logger.Info("successfully restored permission", slog.String("code", restoredPerm.Code), slog.Time("restored_at", time.Now()))
	return restoredPerm, nil
}
