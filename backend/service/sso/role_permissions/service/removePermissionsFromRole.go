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

func (s *Service) RemovePermissionsFromRole(ctx context.Context, clientID uuid.UUID, roleID uuid.UUID, appID int, permissionIDs []uuid.UUID) (*models.OperationStatus, error) {
	const op = "service.RolePermissions.RemovePermissionsFromRole"
	logger := s.logger.With(slog.String("op", op), slog.String("clientID", clientID.String()), slog.String("roleID", roleID.String()), slog.Int("app_id", appID), slog.Int("permissionIDs", len(permissionIDs)))
	logger.Debug("attempting remove permissions to role")

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

	role, err := s.roleProvider.GetByID(ctx, clientID, roleID, appID)
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

	if len(permissionIDs) == 0 {
		logger.Warn("empty permissions list")
		return nil, fmt.Errorf("%w: permission_ids", ErrInvalidArgument)
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

	err := s.relProvider.RemoveRolePermissions(ctx, roleID, clientID, appID, permissionIDs)
	if err != nil {
		logger.Error("failed to remove permissions", slog.Any("error", err), slog.Int("attempted_count", len(permissionIDs)))
		return nil, ErrInternal
	}

	logger.Info("permissions removed", slog.Int("requested_count", len(permissionIDs)))
	return &models.OperationStatus{
		Success:       true,
		Message:       fmt.Sprintf("Removed %d/%d permissions", removedCount, len(permissionIDs)),
		OperationTime: time.Now(),
	}, nil
}
