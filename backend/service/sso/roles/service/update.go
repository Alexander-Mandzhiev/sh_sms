package service

import (
	"backend/service/constants"
	"backend/service/sso/models"
	"backend/service/utils"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log/slog"
	"time"
)

func (s *Service) Update(ctx context.Context, updateData *models.Role) (*models.Role, error) {
	const op = "service.Roles.Update"
	logger := s.logger.With(slog.String("op", op), slog.String("client_id", updateData.ClientID.String()), slog.String("role_id", updateData.ID.String()))
	logger.Debug("attempting to update role")

	if updateData.ClientID == uuid.Nil {
		logger.Warn("empty client_id")
		return nil, fmt.Errorf("%w: client_id", constants.ErrInvalidArgument)
	}
	if updateData.ID == uuid.Nil {
		logger.Warn("empty role_id")
		return nil, fmt.Errorf("%w: role_id", constants.ErrInvalidArgument)
	}

	currentRole, err := s.provider.GetByID(ctx, updateData.ClientID, updateData.ID)
	if err != nil {
		if errors.Is(err, constants.ErrNotFound) {
			logger.Warn("role not found")
			return nil, fmt.Errorf("%w: role", constants.ErrNotFound)
		}
		logger.Error("fetch role failed", slog.Any("error", err))
		return nil, fmt.Errorf("%w: %v", constants.ErrInternal, err)
	}

	if currentRole.ClientID != updateData.ClientID {
		logger.Warn("client ID mismatch", slog.String("stored_client", currentRole.ClientID.String()), slog.String("request_client", updateData.ClientID.String()))
		return nil, fmt.Errorf("%w: role access denied", constants.ErrPermissionDenied)
	}

	updatedRole := currentRole
	needUpdate := false

	if updateData.Name != "" {
		if err = utils.ValidateRoleName(updateData.Name); err != nil {
			logger.Warn("invalid name", slog.Any("error", err))
			return nil, fmt.Errorf("%w: %v", constants.ErrInvalidArgument, err)
		}
		if updatedRole.Name != updateData.Name {
			var exists bool
			exists, err = s.provider.RoleExists(ctx, updateData.ClientID, updateData.Name)
			if err != nil {
				logger.Error("name check failed", slog.Any("error", err))
				return nil, fmt.Errorf("%w: %v", constants.ErrInternal, err)
			}
			if exists {
				logger.Warn("role name conflict", slog.String("name", updateData.Name))
				return nil, fmt.Errorf("%w: role name", constants.ErrConflict)
			}
			updatedRole.Name = updateData.Name
			needUpdate = true
		}
	}

	if updateData.Description != "" {
		updatedRole.Description = updateData.Description
		needUpdate = true
	}

	if updateData.Level >= 0 && updateData.Level != currentRole.Level {
		if err = utils.ValidateRoleLevel(int32(updateData.Level)); err != nil {
			logger.Warn("invalid level", slog.Int("level", updateData.Level), slog.Any("error", err))
			return nil, fmt.Errorf("%w: %v", constants.ErrInvalidArgument, err)
		}
		updatedRole.Level = updateData.Level
		needUpdate = true
	}

	if updateData.IsActive != currentRole.IsActive {
		if !currentRole.IsCustom && !updateData.IsActive {
			return nil, fmt.Errorf("%w: system roles cannot be deactivated", constants.ErrPermissionDenied)
		}
		updatedRole.IsActive = updateData.IsActive
		if updateData.IsActive {
			updatedRole.DeletedAt = nil
		} else {
			now := time.Now().UTC()
			updatedRole.DeletedAt = &now
		}

		needUpdate = true
	}

	if updateData.IsCustom != currentRole.IsCustom {
		logger.Warn("attempt to change system role flag")
		return nil, fmt.Errorf("%w: is_custom cannot be changed", constants.ErrPermissionDenied)
	}

	if !needUpdate {
		logger.Debug("no changes detected")
		return currentRole, nil
	}

	result, err := s.provider.Update(ctx, updatedRole)
	if err != nil {
		logger.Error("update failed", slog.Any("error", err))
		return nil, fmt.Errorf("%w: %v", constants.ErrInternal, err)
	}

	logger.Info("role updated successfully")
	return result, nil
}
