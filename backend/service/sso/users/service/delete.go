package service

import (
	"backend/service/constants"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log/slog"
)

func (s *Service) Delete(ctx context.Context, clientID, userID uuid.UUID, permanent bool) error {
	const op = "service.User.Delete"
	logger := s.logger.With(slog.String("op", op), slog.String("user_id", userID.String()), slog.String("client_id", clientID.String()), slog.Bool("permanent", permanent))
	logger.Debug("attempting to delete user")

	existingUser, err := s.provider.Get(ctx, clientID, userID)
	if err != nil {
		if errors.Is(err, constants.ErrNotFound) {
			logger.Warn("user not found for deletion")
			return fmt.Errorf("%w: user not found", constants.ErrNotFound)
		}
		logger.Error("failed to fetch user", slog.Any("error", err))
		return fmt.Errorf("%w: database error", constants.ErrInternal)
	}

	if existingUser.ClientID != clientID {
		logger.Warn("client ID mismatch", slog.String("expected", clientID.String()), slog.String("actual", existingUser.ClientID.String()))
		return fmt.Errorf("%w: access denied", constants.ErrPermissionDenied)
	}

	if permanent {
		err = s.provider.HardDeleteUser(ctx, clientID, userID)
		logger.Debug("executing hard delete")
	} else {
		if existingUser.DeletedAt != nil {
			logger.Warn("user already deleted")
			return fmt.Errorf("%w: user already deleted", constants.ErrConflict)
		}
		err = s.provider.SoftDeleteUser(ctx, clientID, userID)
		logger.Debug("executing soft delete")
	}

	if err != nil {
		logger.Error("delete operation failed", slog.Any("error", err), slog.Bool("permanent", permanent))
		if errors.Is(err, constants.ErrUserNotFound) {
			return fmt.Errorf("%w: user not found", constants.ErrUserNotFound)
		}
		return fmt.Errorf("%w: delete operation failed", constants.ErrInternal)
	}

	action := "soft deleted"
	if permanent {
		action = "permanently deleted"
	}
	logger.Info("user " + action + " successfully")
	return nil
}
