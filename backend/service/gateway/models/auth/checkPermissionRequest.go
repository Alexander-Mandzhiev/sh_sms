package auth_models

import (
	"github.com/google/uuid"
)

type CheckPermissionRequest struct {
	ClientID   uuid.UUID `json:"client_id" binding:"required,uuid4"`
	AppID      int       `json:"app_id" binding:"required,gt=0"`
	Permission string    `json:"permission" binding:"required"`
	Resource   string    `json:"resource,omitempty"`
}
