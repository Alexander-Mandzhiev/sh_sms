package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID           uuid.UUID
	ClientID     uuid.UUID
	Email        string
	PasswordHash string
	FullName     string
	Phone        string
	IsActive     bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}
