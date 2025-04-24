package service

import (
	"backend/service/constants"
	"backend/service/sso/models"
	"backend/service/sso/roles/service"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log/slog"
)

func (s *service.Service) AddPermission(ctx context.Context, clientID, roleID, permissionID uuid.UUID) (*models.Role, error) {
	const op = "service.Roles.AddPermission"
	logger := s.logger.With(slog.String("op", op), slog.String("role_id", roleID.String()), slog.String("permission_id", permissionID.String()), slog.String("client_id", clientID.String()))
	logger.Debug("attempting to add permission to role")

	role, err := s.provider.GetByID(ctx, clientID, roleID)
	if err != nil {
		if errors.Is(err, constants.ErrNotFound) {
			logger.Warn("role not found")
			return nil, fmt.Errorf("%w: role", constants.ErrNotFound)
		}
		logger.Error("failed to get role", slog.Any("error", err))
		return nil, fmt.Errorf("%w: %v", constants.ErrInternal, err)
	}

	permission, err := s.provider.GetPermissionByID(ctx, permissionID)
	if err != nil {
		if errors.Is(err, constants.ErrNotFound) {
			logger.Warn("permission not found")
			return nil, fmt.Errorf("%w: permission", constants.ErrNotFound)
		}
		logger.Error("failed to get permission", slog.Any("error", err))
		return nil, fmt.Errorf("%w: %v", constants.ErrInternal, err)
	}

	if permission.ClientID != role.ClientID {
		logger.Warn("client ID mismatch", slog.String("role_client", role.ClientID.String()), slog.String("perm_client", permission.ClientID.String()))
		return nil, fmt.Errorf("%w: cross-client operation", constants.ErrPermissionDenied)
	}

	if permission.Level > role.Level {
		logger.Warn("permission level exceeds role level", slog.Int("perm_level", permission.Level), slog.Int("role_level", role.Level))
		return nil, fmt.Errorf("%w: permission level too high", constants.ErrInvalidRoleLevel)
	}

	exists, err := s.provider.PermissionExists(ctx, roleID, permissionID)
	if err != nil {
		logger.Error("permission check failed", slog.Any("error", err))
		return nil, fmt.Errorf("%w: %v", constants.ErrInternal, err)
	}
	if exists {
		logger.Warn("permission already added")
		return nil, fmt.Errorf("%w: permission exists", constants.ErrConflict)
	}

	if err = s.provider.AddPermissionToRole(ctx, roleID, permissionID); err != nil {
		logger.Error("database operation failed", slog.Any("error", err))
		return nil, fmt.Errorf("%w: %v", constants.ErrInternal, err)
	}

	updatedRole, err := s.provider.GetByID(ctx, clientID, roleID)
	if err != nil {
		logger.Error("failed to fetch updated role", slog.Any("error", err))
		return nil, fmt.Errorf("%w: %v", constants.ErrInternal, err)
	}

	logger.Info("permission added successfully")
	return updatedRole, nil
}
