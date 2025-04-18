package models

import (
	"github.com/google/uuid"
	"time"
)

type Role struct {
	ID           uuid.UUID
	ClientID     uuid.UUID
	Name         string
	Description  string
	Level        int
	ParentRoleID *uuid.UUID
	IsCustom     bool
	CreatedBy    *uuid.UUID
	CreatedAt    time.Time
	DeletedAt    *time.Time
	Permissions  []Permission
}
