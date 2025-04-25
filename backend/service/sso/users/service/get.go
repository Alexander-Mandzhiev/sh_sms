package service

import (
	"backend/service/sso/models"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log/slog"
)

func (s *Service) Get(ctx context.Context, clientID, userID uuid.UUID) (*models.User, error) {
	const op = "service.User.Get"
	logger := s.logger.With(slog.String("op", op), slog.String("user_id", userID.String()), slog.String("client_id", clientID.String()))
	logger.Debug("attempting to get user")

	user, err := s.provider.GetByID(ctx, clientID, userID)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			logger.Warn("user not found")
			return nil, fmt.Errorf("%w: %v", ErrNotFound, err)
		}
		logger.Error("database error", slog.Any("error", err))
		return nil, fmt.Errorf("%w: failed to get user", ErrInternal)
	}

	if user.ClientID != clientID {
		logger.Warn("user client ID mismatch", slog.String("expected_client", clientID.String()), slog.String("actual_client", user.ClientID.String()))
		return nil, fmt.Errorf("%w: user doesn't belong to client", ErrPermissionDenied)
	}

	logger.Debug("user retrieved successfully")
	return user, nil
}
