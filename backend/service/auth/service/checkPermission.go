package service

import (
	"backend/pkg/jwt_manager"
	"backend/protos/gen/go/apps/clients_apps"
	"backend/protos/gen/go/sso/role_permissions"
	"backend/service/auth/handle"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log/slog"
)

func (s *AuthService) CheckPermission(ctx context.Context, clientID uuid.UUID, appID int, resource, token, permission string) (bool, []string, []string, error) {
	const op = "auth.service.CheckPermission"
	logger := s.logger.With(slog.String("op", op), slog.String("client_id", clientID.String()), slog.Int("app_id", appID))
	logger.Debug("starting permission check")

	clientApp, err := s.clientApps.GetClientApp(ctx, &client_apps.IdentifierRequest{ClientId: clientID.String(), AppId: int32(appID)})
	if err != nil {
		logger.Error("client/app check failed", slog.Any("error", err))
		return false, nil, nil, fmt.Errorf("%s: %w", op, err)
	}
	if clientApp == nil || !clientApp.IsActive {
		logger.Warn("invalid or inactive client/app", "client_id", clientID.String(), "app_id", appID, "is_active", clientApp.GetIsActive())
		return false, nil, nil, errors.New("invalid client or application")
	}

	secret, err := s.getJWTSecret(ctx, clientID, appID, jwt_manager.AccessToken)
	if err != nil {
		logger.Error("failed to get JWT secret", slog.Any("error", err), slog.String("token_type", "refresh"))
		return false, nil, nil, fmt.Errorf("secret retrieval failed: %w", err)
	}
	logger.Debug("JWT secret successfully retrieved")

	claims, err := jwt_manager.Parse(token, secret, jwt_manager.AccessToken)
	if err != nil {
		logger.Warn("failed to parse refresh token", slog.Any("error", err), slog.String("token_hash", jwt_manager.HashToken(token)))
		return false, nil, nil, fmt.Errorf("%w: invalid token", handle.ErrInvalidToken)
	}
	logger.Debug("token parsed successfully", slog.String("user_id", claims.UserID.String()))

	if claims.ClientID != clientID || claims.AppID != appID {
		logger.Warn("client/app mismatch", slog.String("token_client", claims.ClientID.String()), slog.Int("token_app", claims.AppID))
		return false, nil, nil, handle.ErrPermissionDenied
	}

	logger.Debug("fetching user roles")
	userRoles, err := s.getUserRoles(ctx, claims.UserID, claims.ClientID, claims.AppID)
	if err != nil {
		logger.Error("failed to get user roles", slog.Any("error", err))
		return false, nil, nil, fmt.Errorf("%s: %w", op, err)
	}
	logger.Debug("user roles retrieved", "roles", userRoles)

	logger.Debug("fetching user permissions")
	userPermissions, err := s.getUserPermissions(ctx, claims.ClientID, claims.AppID, userRoles)
	if err != nil {
		logger.Error("failed to get user permissions", slog.Any("error", err))
		return false, nil, nil, fmt.Errorf("%s: %w", op, err)
	}
	logger.Debug("user permissions retrieved", "permissions", userPermissions)

	for _, p := range userPermissions {
		if p == permission {
			logger.Debug("permission granted")
			return true, nil, nil, nil
		}
	}

	rolesResp, err := s.rolePermission.ListRolesForPermission(ctx, &role_permissions.ListRolesRequest{
		PermissionId: permission,
		ClientId:     clientID.String(),
		AppId:        int32(appID),
	})
	if err != nil {
		logger.Error("failed to get required roles", slog.Any("error", err))
		return false, nil, nil, fmt.Errorf("%s: %w", op, err)
	}

	existingRoles := make(map[string]bool)
	for _, role := range userRoles {
		existingRoles[role] = true
	}

	var missingRoles []string
	for _, roleID := range rolesResp.GetRoleIds() {
		if !existingRoles[roleID] {
			missingRoles = append(missingRoles, roleID)
		}
	}

	missingPerms := []string{permission}
	logger.Info("permission check result", slog.Bool("allowed", false), slog.Int("missing_roles", len(missingRoles)), slog.Int("missing_permissions", len(missingPerms)))
	return false, missingRoles, missingPerms, nil
}
