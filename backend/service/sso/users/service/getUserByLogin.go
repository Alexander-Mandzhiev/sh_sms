package service

import (
	"backend/service/sso/models"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
)

func (s *Service) GetUserByLogin(ctx context.Context, login, password string, clientID uuid.UUID) (*models.UserInfo, error) {
	const op = "service.GetUserByLogin"
	logger := s.logger.With(slog.String("op", op), slog.String("email", login), slog.String("client_id", clientID.String()))

	logger.Debug("attempting to authenticate user")

	user, err := s.provider.GetByEmailOrPhone(ctx, clientID, login)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			logger.Warn("user not found")
			return nil, fmt.Errorf("%s: %w", op, ErrNotFound)
		}
		logger.Error("failed to get user by email", slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", op, ErrInternal)
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		logger.Warn("invalid password")
		return nil, fmt.Errorf("%s: %w", op, ErrPermissionDenied)
	}

	logger.Debug("user authenticated successfully", slog.String("user_id", user.ID.String()))

	return &models.UserInfo{
		ID:       user.ID,
		Email:    user.Email,
		Phone:    user.Phone,
		FullName: user.FullName,
		IsActive: user.IsActive,
	}, nil
}
