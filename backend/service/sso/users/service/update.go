package service

import (
	"backend/pkg/utils"
	"backend/service/sso/models"
	"context"
	"errors"
	"fmt"
	"log/slog"
	"strings"
)

func (s *Service) Update(ctx context.Context, updateData *models.User) (*models.User, error) {
	const op = "service.User.Update"
	logger := s.logger.With(slog.String("op", op), slog.String("user_id", updateData.ID.String()), slog.String("client_id", updateData.ClientID.String()))
	logger.Debug("attempting to update user")

	existingUser, err := s.provider.GetByID(ctx, updateData.ClientID, updateData.ID)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			logger.Warn("user not found for update")
			return nil, fmt.Errorf("%w: user not found", ErrNotFound)
		}
		logger.Error("failed to fetch user", slog.Any("error", err))
		return nil, fmt.Errorf("%w: database error", ErrInternal)
	}

	if existingUser.ClientID != updateData.ClientID {
		logger.Warn("client ID mismatch")
		return nil, fmt.Errorf("%w: access denied", ErrPermissionDenied)
	}

	updatedUser := *existingUser
	hasChanges := false

	if updateData.Email != "" && updateData.Email != existingUser.Email {
		if err = utils.ValidateEmail(updateData.Email); err != nil {
			logger.Warn("invalid email format", slog.Any("error", err))
			return nil, fmt.Errorf("%w: %v", ErrInvalidArgument, err)
		}

		var conflictUser *models.User
		conflictUser, err = s.provider.GetByEmail(ctx, updateData.ClientID, updateData.Email)
		if err != nil && !errors.Is(err, ErrNotFound) {
			logger.Error("email check failed", slog.Any("error", err))
			return nil, fmt.Errorf("%w: email check failed", ErrInternal)
		}

		if conflictUser != nil && conflictUser.ID != existingUser.ID {
			logger.Warn("email already exists", slog.String("email", updateData.Email))
			return nil, fmt.Errorf("%w: email already registered", ErrConflict)
		}

		updatedUser.Email = updateData.Email
		hasChanges = true
	}

	if updateData.FullName != existingUser.FullName {
		trimmed := strings.TrimSpace(updateData.FullName)
		if trimmed == "" {
			logger.Warn("empty full name")
			return nil, fmt.Errorf("%w: full_name cannot be empty", ErrInvalidArgument)
		}
		updatedUser.FullName = trimmed
		hasChanges = true
	}

	if updateData.Phone != existingUser.Phone {
		if updateData.Phone != "" {
			if err = utils.ValidatePhone(updateData.Phone); err != nil {
				logger.Warn("invalid phone format", slog.Any("error", err))
				return nil, fmt.Errorf("%w: %v", ErrInvalidArgument, err)
			}
		}
		updatedUser.Phone = updateData.Phone
		hasChanges = true
	}

	if !hasChanges {
		logger.Debug("no changes detected")
		return existingUser, nil
	}

	if err = s.provider.Update(ctx, &updatedUser); err != nil {
		logger.Error("update operation failed", slog.Any("error", err))
		if errors.Is(err, ErrConflict) {
			return nil, fmt.Errorf("%w: data conflict", ErrConflict)
		}
		return nil, fmt.Errorf("%w: failed to update user", ErrInternal)
	}

	logger.Info("user updated successfully")
	return &updatedUser, nil
}
