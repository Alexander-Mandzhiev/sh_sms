package service

import (
	"backend/service/constants"
	"backend/service/sso/models"
	"backend/service/utils"
	"context"
	"errors"
	"fmt"
	"log/slog"
	"strings"
	"time"
)

func (s *Service) Update(ctx context.Context, updateData *models.User) (*models.User, error) {
	const op = "service.User.Update"
	logger := s.logger.With(slog.String("op", op), slog.String("user_id", updateData.ID.String()), slog.String("client_id", updateData.ClientID.String()))
	logger.Debug("attempting to update user")

	existingUser, err := s.provider.Get(ctx, updateData.ClientID, updateData.ID)
	if err != nil {
		if errors.Is(err, constants.ErrNotFound) {
			logger.Warn("user not found for update")
			return nil, fmt.Errorf("%w: user not found", constants.ErrNotFound)
		}
		logger.Error("failed to fetch user", slog.Any("error", err))
		return nil, fmt.Errorf("%w: database error", constants.ErrInternal)
	}

	if existingUser.ClientID != updateData.ClientID {
		logger.Warn("client ID mismatch", slog.String("expected", updateData.ClientID.String()), slog.String("actual", existingUser.ClientID.String()))
		return nil, fmt.Errorf("%w: access denied", constants.ErrPermissionDenied)
	}

	partialUpdate := models.UserUpdate{
		UpdatedAt: time.Now().UTC(),
	}

	if updateData.Email != "" && updateData.Email != existingUser.Email {
		var conflictUser *models.User
		if err = utils.ValidateEmail(updateData.Email); err != nil {
			logger.Warn("invalid email format", slog.Any("error", err))
			return nil, fmt.Errorf("%w: %v", constants.ErrInvalidArgument, err)
		}

		conflictUser, err = s.provider.GetByEmail(ctx, updateData.ClientID, updateData.Email)
		if err != nil && !errors.Is(err, constants.ErrNotFound) {
			logger.Error("email check failed", slog.Any("error", err))
			return nil, fmt.Errorf("%w: email check failed", constants.ErrInternal)
		}

		if conflictUser != nil && conflictUser.ID != existingUser.ID {
			logger.Warn("email already exists", slog.String("email", updateData.Email))
			return nil, fmt.Errorf("%w: email already registered", constants.ErrConflict)
		}
		partialUpdate.Email = &updateData.Email
	}

	if updateData.FullName != existingUser.FullName {
		trimmed := strings.TrimSpace(updateData.FullName)
		if trimmed == "" {
			logger.Warn("empty full name")
			return nil, fmt.Errorf("%w: full_name cannot be empty", constants.ErrInvalidArgument)
		}
		partialUpdate.FullName = &trimmed
	}

	if updateData.Phone != existingUser.Phone {
		if updateData.Phone != "" {
			if err = utils.ValidatePhone(updateData.Phone); err != nil {
				logger.Warn("invalid phone format", slog.Any("error", err))
				return nil, fmt.Errorf("%w: %v", constants.ErrInvalidArgument, err)
			}
		}
		partialUpdate.Phone = &updateData.Phone
	}

	if updateData.IsActive != existingUser.IsActive {
		partialUpdate.IsActive = &updateData.IsActive
	}

	if partialUpdate.Email == nil && partialUpdate.FullName == nil && partialUpdate.Phone == nil && partialUpdate.IsActive == nil {
		logger.Debug("no changes detected")
		return existingUser, nil
	}

	if err = s.provider.Update(ctx, existingUser.ID, updateData.ClientID, partialUpdate); err != nil {
		logger.Error("update operation failed", slog.Any("error", err))
		if errors.Is(err, constants.ErrConflict) {
			return nil, fmt.Errorf("%w: data conflict", constants.ErrConflict)
		}
		return nil, fmt.Errorf("%w: failed to update user", constants.ErrInternal)
	}

	if partialUpdate.Email != nil {
		existingUser.Email = *partialUpdate.Email
	}
	if partialUpdate.FullName != nil {
		existingUser.FullName = *partialUpdate.FullName
	}
	if partialUpdate.Phone != nil {
		existingUser.Phone = *partialUpdate.Phone
	}
	if partialUpdate.IsActive != nil {
		existingUser.IsActive = *partialUpdate.IsActive
	}
	existingUser.UpdatedAt = partialUpdate.UpdatedAt

	logger.Info("user updated successfully")
	return existingUser, nil
}
