package sso_models

import "github.com/google/uuid"

type RolePermissionsRequest struct {
	RoleId        uuid.UUID `json:"role_id"`
	ClientID      string    `json:"client_id" binding:"required,uuid"`
	AppID         int32     `json:"app_id" binding:"required,min=1"`
	PermissionIDs []string  `json:"permission_ids" binding:"required,min=1"`
}
