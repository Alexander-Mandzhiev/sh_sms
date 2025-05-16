package service

import (
	"backend/pkg/utils"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log/slog"
)

func (s *Service) Delete(ctx context.Context, clientID, roleID uuid.UUID, appID int, permanent bool) error {
	const op = "service.Roles.Delete"
	logger := s.logger.With(slog.String("op", op), slog.String("client_id", clientID.String()), slog.String("role_id", "REDACTED"), slog.Int("app_id", appID), slog.Bool("permanent", permanent))
	logger.Debug("delete operation started")

	if clientID == uuid.Nil {
		logger.Warn("invalid client_id: nil UUID")
		return fmt.Errorf("%w: client_id", ErrInvalidArgument)
	}

	if roleID == uuid.Nil {
		logger.Warn("invalid role_id: nil UUID")
		return fmt.Errorf("%w: role_id", ErrInvalidArgument)
	}

	if err := utils.ValidateAppID(appID); err != nil {
		logger.Warn("invalid app_id", slog.Int("value", appID))
		return fmt.Errorf("%w: app_id", ErrInvalidArgument)
	}

	err := s.provider.Delete(ctx, clientID, roleID, appID, permanent)
	if err != nil {
		logMsg := "delete operation failed"
		if errors.Is(err, ErrNotFound) {
			logger.Warn(logMsg, slog.String("error_type", "not_found"))
			return fmt.Errorf("%w: %v", ErrNotFound, "role not found")
		}
		if errors.Is(err, ErrConflict) {
			logger.Warn(logMsg, slog.String("error_type", "conflict"))
			return fmt.Errorf("%w: %v", ErrConflict, err)
		}
		logger.Error(logMsg, slog.String("error", err.Error()))
		return fmt.Errorf("%w: %v", ErrInternal, "internal server error")
	}

	deleteType := "soft"
	if permanent {
		deleteType = "hard"
	}

	logger.Info("role deleted successfully", slog.String("type", deleteType), slog.Int("app_id", appID))
	return nil
}
