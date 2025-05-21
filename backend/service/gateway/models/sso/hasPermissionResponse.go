package sso_models

import (
	"backend/protos/gen/go/sso/role_permissions"
	"time"
)

type HasPermissionResponse struct {
	HasPermission bool      `json:"has_permission"`
	CheckedAt     time.Time `json:"checked_at"`
}

func HasPermissionResponseFromProto(pb *role_permissions.HasPermissionResponse) *HasPermissionResponse {
	return &HasPermissionResponse{
		HasPermission: pb.HasPermission,
		CheckedAt:     pb.CheckedAt.AsTime(),
	}
}
