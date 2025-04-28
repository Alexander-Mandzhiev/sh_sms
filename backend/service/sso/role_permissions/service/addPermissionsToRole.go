package service

import (
	"backend/service/sso/models"
	servicePermission "backend/service/sso/permissions/service"
	serviceRole "backend/service/sso/roles/service"
	"backend/service/utils"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log/slog"
	"time"
)

func (s *Service) AddPermissionsToRole(ctx context.Context, clientID uuid.UUID, roleID uuid.UUID, appID int, permissionIDs []uuid.UUID) (*models.OperationStatus, error) {
	const op = "service.RolePermissions.AddPermissionsToRole"
	logger := s.logger.With(slog.String("op", op), slog.String("clientID", clientID.String()), slog.String("roleID", roleID.String()), slog.Int("app_id", appID), slog.Int("permissionIDs", len(permissionIDs)))
	logger.Debug("attempting add permissions to role")

	if err := utils.ValidateUUID(clientID); err != nil {
		logger.Warn("invalid client_id", slog.Any("error", err))
		return nil, fmt.Errorf("%w: client_id", ErrInvalidArgument)
	}

	if err := utils.ValidateUUID(roleID); err != nil {
		logger.Warn("invalid role_id", slog.Any("error", err))
		return nil, fmt.Errorf("%w: role_id", ErrInvalidArgument)
	}

	if err := utils.ValidateAppID(appID); err != nil {
		logger.Warn("invalid app_id", slog.Any("error", err))
		return nil, fmt.Errorf("%w: app_id", ErrInvalidArgument)
	}

	role, err := s.roleProvider.GetByID(ctx, clientID, roleID)
	if err != nil {
		if errors.Is(err, serviceRole.ErrNotFound) {
			return nil, ErrRoleNotFound
		}
		logger.Error("failed to get role", slog.Any("error", err))
		return nil, ErrInternal
	}

	if !role.IsActive {
		return nil, fmt.Errorf("%w: role is inactive", ErrInactiveEntity)
	}

	for _, permID := range permissionIDs {
		if err = utils.ValidateUUID(permID); err != nil {
			logger.Warn("invalid permission_id", slog.Any("error", err))
			return nil, fmt.Errorf("%w: permission_id", ErrInvalidArgument)
		}
		var perm *models.Permission
		perm, err = s.permProvider.GetByID(ctx, permID, appID)
		if err != nil {
			if errors.Is(err, servicePermission.ErrNotFound) {
				return nil, ErrPermissionNotFound
			}
			logger.Error("failed to get permission", slog.Any("error", err))
			return nil, ErrInternal
		}

		if !perm.IsActive {
			return nil, fmt.Errorf("%w: permission %s is inactive", ErrInactiveEntity, permID)
		}
	}

	if err = s.relProvider.AddRolePermissions(ctx, roleID, clientID, appID, permissionIDs); err != nil {
		logger.Error("failed to add permissions", slog.Any("error", err))
		return nil, ErrInternal
	}

	return &models.OperationStatus{
		Success:       true,
		Message:       fmt.Sprintf("Successfully added %d permissions", len(permissionIDs)),
		OperationTime: time.Now(),
	}, nil
}
