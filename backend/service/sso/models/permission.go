package models

import (
	"github.com/google/uuid"
	"time"
)

type Permission struct {
	ID          uuid.UUID
	Code        string
	Description string
	Category    string
	AppID       int
	CreatedAt   time.Time
	DeletedAt   *time.Time
}
