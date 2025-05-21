package sso_models

import "backend/protos/gen/go/sso/role_permissions"

type RolePermissionsList struct {
	PermissionIDs []string `json:"permission_ids"`
}

func RolePermissionsListFromProto(pb *role_permissions.ListPermissionsResponse) *RolePermissionsList {
	return &RolePermissionsList{
		PermissionIDs: pb.PermissionIds,
	}
}
