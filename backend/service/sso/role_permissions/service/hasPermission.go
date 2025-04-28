package service

import (
	servicePermission "backend/service/sso/permissions/service"
	serviceRole "backend/service/sso/roles/service"
	"backend/service/utils"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log/slog"
)

func (s *Service) HasPermission(ctx context.Context, clientID uuid.UUID, roleID uuid.UUID, permissionID uuid.UUID, appID int) (bool, error) {
	const op = "service.HasPermission"
	logger := s.logger.With(slog.String("op", op), slog.String("clientID", clientID.String()), slog.String("roleID", roleID.String()), slog.String("permissionID", permissionID.String()), slog.Int("appID", appID))
	logger.Debug("starting permission check")

	if err := utils.ValidateUUID(clientID); err != nil {
		logger.Warn("invalid client_id format", slog.Any("error", err))
		return false, fmt.Errorf("%w: client_id", ErrInvalidArgument)
	}

	if err := utils.ValidateUUID(roleID); err != nil {
		logger.Warn("invalid role_id format", slog.Any("error", err))
		return false, fmt.Errorf("%w: role_id", ErrInvalidArgument)
	}

	if err := utils.ValidateUUID(permissionID); err != nil {
		logger.Warn("invalid permission_id format", slog.Any("error", err))
		return false, fmt.Errorf("%w: permission_id", ErrInvalidArgument)
	}

	if err := utils.ValidateAppID(appID); err != nil {
		logger.Warn("invalid app_id", slog.Any("error", err))
		return false, fmt.Errorf("%w: app_id", ErrInvalidArgument)
	}

	role, err := s.roleProvider.GetByID(ctx, clientID, roleID)
	if err != nil {
		if errors.Is(err, serviceRole.ErrNotFound) {
			logger.Warn("role not found")
			return false, ErrRoleNotFound
		}
		logger.Error("failed to fetch role", slog.Any("error", err))
		return false, ErrInternal
	}

	if !role.IsActive {
		logger.Info("role is inactive, automatic denial")
		return false, nil
	}

	perm, err := s.permProvider.GetByID(ctx, permissionID, appID)
	if err != nil {
		if errors.Is(err, servicePermission.ErrNotFound) {
			logger.Warn("permission not found")
			return false, ErrPermissionNotFound
		}
		logger.Error("failed to fetch permission", slog.Any("error", err))
		return false, ErrInternal
	}

	if !perm.IsActive {
		logger.Info("permission is inactive, automatic denial")
		return false, nil
	}

	hasRelation, err := s.relProvider.HasRelation(ctx, roleID, permissionID)
	if err != nil {
		logger.Error("failed to check permission relation", slog.Any("error", err), slog.String("relation_check", "role_permission"))
		return false, ErrInternal
	}

	logger.Debug("permission check completed", slog.Bool("has_permission", hasRelation))
	return hasRelation, nil
}
