package sso_models

import "github.com/google/uuid"

type HasPermissionRequest struct {
	RoleID       uuid.UUID `json:"role_id" binding:"required,uuid"`
	ClientID     uuid.UUID `json:"client_id" binding:"required,uuid"`
	PermissionID uuid.UUID `json:"permission_id" binding:"required,uuid"`
	AppID        int32     `json:"app_id" binding:"required,min=1"`
}
