package service

import (
	"backend/protos/gen/go/sso/role_permissions"
	"context"
	"fmt"
	"github.com/google/uuid"
)

func (s *AuthService) getUserPermissions(ctx context.Context, clientID uuid.UUID, appID int, roles []string) ([]string, error) {
	permissionsSet := make(map[string]struct{})
	for _, roleID := range roles {
		permissions, err := s.rolePermission.ListPermissionsForRole(ctx, &role_permissions.ListPermissionsRequest{
			RoleId:   roleID,
			ClientId: clientID.String(),
			AppId:    int32(appID),
		})
		if err != nil {
			return nil, fmt.Errorf("failed to get permissions for role %s: %w", roleID, err)
		}

		for _, permID := range permissions.PermissionIds {
			if permID != "" {
				permissionsSet[permID] = struct{}{}
			}
		}
	}

	permissions := make([]string, 0, len(permissionsSet))
	for perm := range permissionsSet {
		permissions = append(permissions, perm)
	}

	return permissions, nil
}
