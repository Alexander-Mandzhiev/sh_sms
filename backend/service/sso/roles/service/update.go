package service

import (
	"backend/service/sso/models"
	"backend/service/utils"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log/slog"
)

func (s *Service) Update(ctx context.Context, updateData *models.Role) (*models.Role, error) {
	const op = "service.Roles.Update"
	logger := s.logger.With(slog.String("op", op), slog.String("client_id", updateData.ClientID.String()), slog.String("role_id", updateData.ID.String()))
	logger.Debug("attempting to update role")

	if updateData.ClientID == uuid.Nil {
		logger.Warn("empty client_id")
		return nil, fmt.Errorf("%w: client_id", ErrInvalidArgument)
	}
	if updateData.ID == uuid.Nil {
		logger.Warn("empty role_id")
		return nil, fmt.Errorf("%w: role_id", ErrInvalidArgument)
	}

	if updateData.AppID <= 0 {
		logger.Warn("invalid app_id")
		return nil, fmt.Errorf("%w: invalid app_id", ErrInvalidArgument)
	}

	currentRole, err := s.provider.GetByID(ctx, updateData.ClientID, updateData.ID, updateData.AppID)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			logger.Warn("role not found")
			return nil, fmt.Errorf("%w: role", ErrNotFound)
		}
		logger.Error("fetch role failed", slog.Any("error", err))
		return nil, fmt.Errorf("%w: %v", ErrInternal, err)
	}

	if currentRole.ClientID != updateData.ClientID {
		logger.Warn("client ID mismatch", slog.String("stored_client", currentRole.ClientID.String()), slog.String("request_client", updateData.ClientID.String()))
		return nil, fmt.Errorf("%w: role access denied", ErrPermissionDenied)
	}

	updatedRole := currentRole
	needUpdate := false

	if updateData.Name != "" {
		if err = utils.ValidateRoleName(updateData.Name, 150); err != nil {
			logger.Warn("invalid name", slog.Any("error", err))
			return nil, fmt.Errorf("%w: %v", ErrInvalidArgument, err)
		}
		if updatedRole.Name != updateData.Name {
			var exists bool
			exists, err = s.provider.RoleExists(ctx, updateData.ClientID, updateData.AppID, updateData.Name)
			if err != nil {
				logger.Error("name check failed", slog.Any("error", err))
				return nil, fmt.Errorf("%w: %v", ErrInternal, err)
			}
			if exists {
				logger.Warn("role name conflict", slog.String("name", updateData.Name))
				return nil, fmt.Errorf("%w: role name", ErrConflict)
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
		if err = utils.ValidateRoleLevel(updateData.Level); err != nil {
			logger.Warn("invalid level", slog.Int("level", updateData.Level), slog.Any("error", err))
			return nil, fmt.Errorf("%w: %v", ErrInvalidArgument, err)
		}
		updatedRole.Level = updateData.Level
		needUpdate = true
	}

	if !currentRole.IsCustom {
		logger.Warn("attempt to change system role flag")
		return nil, fmt.Errorf("%w: is_custom cannot be changed", ErrPermissionDenied)
	}

	if !needUpdate {
		logger.Debug("no changes detected")
		return currentRole, nil
	}

	result, err := s.provider.Update(ctx, updatedRole)
	if err != nil {
		logger.Error("update failed", slog.Any("error", err))
		return nil, fmt.Errorf("%w: %v", ErrInternal, err)
	}

	logger.Info("role updated successfully")
	return result, nil
}
