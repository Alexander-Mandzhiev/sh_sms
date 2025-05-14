package models

import "github.com/google/uuid"

type ListRolesForPermissionRequest struct {
	PermissionID uuid.UUID `json:"permission_id" binding:"required,uuid"`
	ClientID     uuid.UUID `json:"client_id" binding:"required,uuid"`
	AppID        int32     `json:"app_id" binding:"required,min=1"`
}
