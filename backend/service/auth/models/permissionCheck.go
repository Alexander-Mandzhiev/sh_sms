package models

import (
	"github.com/google/uuid"
)

type PermissionCheck struct {
	UserID     uuid.UUID
	ClientID   uuid.UUID
	AppID      int
	Permission string
	Resource   string
}
