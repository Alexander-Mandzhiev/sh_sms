package models

import (
	"github.com/google/uuid"
	"time"
)

type Role struct {
	ID          uuid.UUID
	ClientID    uuid.UUID
	AppID       int
	Name        string
	Description string
	Level       int
	IsCustom    bool
	IsActive    bool
	CreatedBy   *uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
	Permissions []Permission
}
