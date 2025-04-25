package service

import (
	"backend/service/sso/models"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log/slog"
)

func (s *Service) Restore(ctx context.Context, clientID, userID uuid.UUID) (*models.User, error) {
	const op = "service.User.Restore"
	logger := s.logger.With(slog.String("op", op), slog.String("user_id", userID.String()), slog.String("client_id", clientID.String()))
	logger.Debug("attempting to restore user")

	exists, err := s.provider.Exists(ctx, clientID, userID)
	if err != nil {
		logger.Error("existence check failed", slog.Any("error", err))
		return nil, fmt.Errorf("%w: %v", ErrInternal, err)
	}
	if !exists {
		logger.Warn("user not found")
		return nil, fmt.Errorf("%w: user not found", ErrNotFound)
	}

	user, err := s.provider.GetByID(ctx, clientID, userID)
	if err != nil {
		logger.Error("failed to fetch user", slog.Any("error", err))
		return nil, fmt.Errorf("%w: database error", ErrInternal)
	}

	if user.DeletedAt == nil {
		logger.Warn("user is not deleted")
		return nil, fmt.Errorf("%w: user is active", ErrConflict)
	}

	var conflictUser *models.User
	if conflictUser, err = s.provider.GetByEmail(ctx, clientID, user.Email); err != nil {
		if !errors.Is(err, ErrNotFound) {
			logger.Error("email check failed", slog.Any("error", err))
			return nil, fmt.Errorf("%w: email check failed", ErrInternal)
		}
	} else if conflictUser != nil && conflictUser.ID != user.ID {
		logger.Warn("email already exists", slog.String("email", user.Email))
		return nil, fmt.Errorf("%w: email already registered", ErrConflict)
	}

	if user, err = s.provider.Restore(ctx, clientID, userID); err != nil {
		logger.Error("restore failed", slog.Any("error", err))
		return nil, fmt.Errorf("%w: restore operation failed", ErrInternal)
	}

	logger.Info("user restored successfully")
	return user, nil
}
