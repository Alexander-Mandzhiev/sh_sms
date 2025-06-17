package service

import (
	"backend/pkg/utils"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
)

func (s *Service) SetPassword(ctx context.Context, clientID, userID uuid.UUID, password string) error {
	const op = "service.User.SetPassword"
	logger := s.logger.With(slog.String("op", op), slog.String("user_id", userID.String()), slog.String("client_id", clientID.String()))
	logger.Debug("attempting to set password")

	if err := utils.ValidatePassword(password); err != nil {
		logger.Warn("invalid password format", slog.Any("error", err))
		return fmt.Errorf("%w: %v", ErrInvalidArgument, err)
	}

	exist, err := s.provider.Exists(ctx, clientID, userID)
	if err != nil {
		logger.Error("existence check failed", slog.Any("error", err))
		return fmt.Errorf("%w: %v", ErrInternal, err)
	}
	if !exist {
		logger.Warn("user not found")
		return fmt.Errorf("%w: user not found", ErrNotFound)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logger.Error("password hashing failed", slog.Any("error", err))
		return fmt.Errorf("%w: failed to hash password", ErrInternal)
	}

	if err = s.provider.UpdatePasswordHash(ctx, userID, string(hashedPassword)); err != nil {
		logger.Error("failed to update password", slog.Any("error", err))
		if errors.Is(err, ErrNotFound) {
			return fmt.Errorf("%w: user not found", ErrNotFound)
		}
		return fmt.Errorf("%w: failed to update password", ErrInternal)
	}

	logger.Info("password updated successfully")
	return nil
}
