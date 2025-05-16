package service

import (
	"backend/protos/gen/go/sso/roles"
	"context"
	"fmt"
	"github.com/google/uuid"
)

func (s *AuthService) getUserRoles(ctx context.Context, clientID uuid.UUID, appID int) ([]string, error) {
	rolesResp, err := s.roles.ListRoles(ctx, &roles.ListRequest{
		ClientId: clientID.String(),
		AppId:    int32(appID),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get roles: %w", err)
	}

	rolesSet := make([]string, 0, len(rolesResp.Roles))
	for _, role := range rolesResp.Roles {
		if role.Id != "" {
			rolesSet = append(rolesSet, role.Id)
		}
	}
	return rolesSet, nil
}
