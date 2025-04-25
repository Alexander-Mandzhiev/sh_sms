package service

import (
	"backend/service/sso/models"
	"backend/service/utils"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
	"time"
)

func (s *Service) Create(ctx context.Context, user *models.User, password string) error {
	const op = "service.User.Create"
	logger := s.logger.With(slog.String("op", op), slog.String("client_id", user.ClientID.String()))
	logger.Debug("attempting to create user")

	if err := utils.ValidatePasswordPolicy(password); err != nil {
		logger.Warn("password policy validation failed", slog.Any("error", err), slog.Int("password_length", len(password)))
		return fmt.Errorf("%w: %v", ErrInvalidArgument, err)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logger.Error("password hashing failed", slog.Any("error", err))
		return fmt.Errorf("%w: failed to hash password", ErrInternal)
	}

	user.PasswordHash = string(hashedPassword)
	if user.ID == uuid.Nil {
		user.ID = uuid.New()
	}
	user.CreatedAt = time.Now().UTC()
	user.UpdatedAt = user.CreatedAt

	if err = s.provider.Create(ctx, user); err != nil {
		if errors.Is(err, ErrEmailAlreadyExists) {
			logger.Warn("duplicate email", slog.String("email", user.Email))
			return fmt.Errorf("%w: %v", ErrEmailAlreadyExists, user.Email)
		}
		logger.Error("database operation failed", slog.Any("error", err))
		return fmt.Errorf("%w: %v", ErrInternal, err)
	}

	logger.Info("user created successfully", slog.String("user_id", user.ID.String()), slog.String("email", user.Email))
	return nil
}
