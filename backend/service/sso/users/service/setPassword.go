package service

import (
	"backend/service/constants"
	"backend/service/utils"
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

	if err := utils.ValidatePasswordPolicy(password); err != nil {
		logger.Warn("invalid password format", slog.Any("error", err))
		return fmt.Errorf("%w: %v", constants.ErrInvalidArgument, err)
	}

	user, err := s.provider.Get(ctx, clientID, userID)
	if err != nil {
		if errors.Is(err, constants.ErrNotFound) {
			logger.Warn("user not found")
			return fmt.Errorf("%w: user not found", constants.ErrNotFound)
		}
		logger.Error("failed to fetch user", slog.Any("error", err))
		return fmt.Errorf("%w: database error", constants.ErrInternal)
	}

	if user.ClientID != clientID {
		logger.Warn("client ID mismatch", slog.String("expected", clientID.String()), slog.String("actual", user.ClientID.String()))
		return fmt.Errorf("%w: access denied", constants.ErrPermissionDenied)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logger.Error("password hashing failed", slog.Any("error", err))
		return fmt.Errorf("%w: failed to hash password", constants.ErrInternal)
	}

	user.PasswordHash = string(hashedPassword)

	if err = s.provider.UpdatePasswordHash(ctx, userID, string(hashedPassword)); err != nil {
		logger.Error("failed to update password", slog.Any("error", err))
		if errors.Is(err, constants.ErrNotFound) {
			return fmt.Errorf("%w: user not found", constants.ErrNotFound)
		}

		return fmt.Errorf("%w: failed to update password", constants.ErrInternal)
	}

	logger.Info("password updated successfully")
	return nil
}
