package service

import (
	"backend/protos/gen/go/apps/clients_apps"
	"backend/protos/gen/go/sso/role_permissions"
	"backend/service/auth/models"
	"context"
	"errors"
	"fmt"
	"log/slog"
)

func (s *AuthService) CheckPermission(ctx context.Context, check *models.PermissionCheck) (bool, []string, []string, error) {
	const op = "auth.service.CheckPermission"
	logger := s.logger.With(slog.String("op", op), slog.String("user_id", check.UserID.String()), slog.String("client_id", check.ClientID.String()), slog.Int("app_id", check.AppID), slog.String("permission", check.Permission))
	logger.Debug("starting permission check")

	clientApp, err := s.clientApps.GetClientApp(ctx, &client_apps.IdentifierRequest{ClientId: check.ClientID.String(), AppId: int32(check.AppID)})
	if err != nil {
		logger.Error("client/app check failed", slog.Any("error", err))
		return false, nil, nil, fmt.Errorf("%s: %w", op, err)
	}
	if clientApp == nil || !clientApp.IsActive {
		logger.Warn("invalid or inactive client/app", "client_id", check.ClientID.String(), "app_id", check.AppID, "is_active", clientApp.GetIsActive())
		return false, nil, nil, errors.New("invalid client or application")
	}

	logger.Debug("fetching user roles")
	userRoles, err := s.getUserRoles(ctx, check.UserID, check.ClientID, check.AppID)
	if err != nil {
		logger.Error("failed to get user roles", slog.Any("error", err))
		return false, nil, nil, fmt.Errorf("%s: %w", op, err)
	}
	logger.Debug("user roles retrieved", "roles", userRoles)

	logger.Debug("fetching user permissions")
	userPermissions, err := s.getUserPermissions(ctx, check.ClientID, check.AppID, userRoles)
	if err != nil {
		logger.Error("failed to get user permissions", slog.Any("error", err))
		return false, nil, nil, fmt.Errorf("%s: %w", op, err)
	}
	logger.Debug("user permissions retrieved", "permissions", userPermissions)

	permissionsSet := make(map[string]struct{})
	for _, perm := range userPermissions {
		permissionsSet[perm] = struct{}{}
	}

	hasPermission := false
	if _, ok := permissionsSet[check.Permission]; ok {
		logger.Debug("permission found in set")
		hasPermission = true
	} else {
		logger.Debug("permission NOT found in set")
	}

	var missingRoles, missingPerms []string
	if !hasPermission {
		logger.Debug("checking required roles for missing permission")
		rolesResp, err := s.rolePermission.ListRolesForPermission(ctx, &role_permissions.ListRolesRequest{
			PermissionId: check.Permission,
			ClientId:     check.ClientID.String(),
			AppId:        int32(check.AppID),
		})
		if err != nil {
			logger.Error("failed to get roles for permission", slog.Any("error", err))
			return false, nil, nil, fmt.Errorf("failed to find roles for permission: %w", err)
		}

		existingRoles := make(map[string]struct{})
		for _, roleID := range userRoles {
			existingRoles[roleID] = struct{}{}
		}

		for _, roleID := range rolesResp.RoleIds {
			if _, exists := existingRoles[roleID]; !exists {
				missingRoles = append(missingRoles, roleID)
			}
		}
		missingPerms = []string{check.Permission}

		logger.Debug("missing requirements calculated", "missing_roles", missingRoles, "missing_permissions", missingPerms)
	}

	logger.Info("permission check result", "allowed", hasPermission, "missing_roles_count", len(missingRoles), "missing_permissions_count", len(missingPerms))
	return hasPermission, missingRoles, missingPerms, nil
}
