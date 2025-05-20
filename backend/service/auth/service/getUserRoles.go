package service

import (
	"backend/protos/gen/go/sso/users_roles"
	"context"
	"fmt"
	"github.com/google/uuid"
)

func (s *AuthService) getUserRoles(ctx context.Context, userID, clientID uuid.UUID, appID int) ([]string, error) {
	rolesResp, err := s.userRole.ListForUser(ctx, &user_roles.UserRequest{
		UserId:   userID.String(),
		ClientId: clientID.String(),
		AppId:    int32(appID),
		Page:     1,
		Count:    1000,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get roles: %w", err)
	}

	rolesSet := make([]string, 0, len(rolesResp.Assignments))
	for _, role := range rolesResp.Assignments {
		if role.RoleId != "" {
			rolesSet = append(rolesSet, role.RoleId)
		}
	}
	return rolesSet, nil
}
