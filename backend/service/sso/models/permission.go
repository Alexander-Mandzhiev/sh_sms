package models

import (
	"github.com/google/uuid"
	"time"
)

type Permission struct {
	ID          uuid.UUID  `db:"id"`
	Code        string     `db:"code"`
	Description string     `db:"description"`
	Category    string     `db:"category"`
	AppID       int        `db:"app_id"`
	IsActive    bool       `db:"is_active"`
	CreatedAt   time.Time  `db:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at"`
	DeletedAt   *time.Time `db:"deleted_at"`
}
