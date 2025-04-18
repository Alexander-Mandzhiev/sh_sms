package models

import (
	"github.com/google/uuid"
	"time"
)

type UserRole struct {
	UserID     uuid.UUID
	RoleID     uuid.UUID
	ClientID   uuid.UUID
	AssignedBy *uuid.UUID
	ExpiresAt  *time.Time
	CreatedAt  time.Time
}
