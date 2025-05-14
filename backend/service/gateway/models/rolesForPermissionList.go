package models

import "backend/protos/gen/go/sso/role_permissions"

type RolesForPermissionList struct {
	RoleIDs []string `json:"role_ids"`
}

func RolesForPermissionListFromProto(pb *role_permissions.ListRolesResponse) *RolesForPermissionList {
	return &RolesForPermissionList{
		RoleIDs: pb.RoleIds,
	}
}
