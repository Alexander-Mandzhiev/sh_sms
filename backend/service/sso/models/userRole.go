package models

import (
	"github.com/google/uuid"
	"time"
)

type UserRole struct {
	UserID     uuid.UUID
	RoleID     uuid.UUID
	ClientID   uuid.UUID
	AppID      int
	AssignedBy uuid.UUID
	ExpiresAt  *time.Time
	AssignedAt time.Time
}
