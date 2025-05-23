package auth_models

import "backend/protos/gen/go/auth"

type CheckPermissionResponse struct {
	Allowed            bool     `json:"allowed" validate:"required"`
	MissingRoles       []string `json:"missing_roles,omitempty" validate:"dive,required"`
	MissingPermissions []string `json:"missing_permissions,omitempty" validate:"dive,required"`
}

func CheckPermissionFromProtoResponse(pb *auth.PermissionCheckResponse) *CheckPermissionResponse {
	if pb == nil {
		return &CheckPermissionResponse{}
	}

	return &CheckPermissionResponse{
		Allowed:            pb.GetAllowed(),
		MissingRoles:       pb.GetMissingRoles(),
		MissingPermissions: pb.GetMissingPermissions(),
	}
}
