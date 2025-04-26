package models

import (
	"github.com/google/uuid"
	"time"
)

type Permission struct {
	ID          uuid.UUID  `json:"id" db:"id"`
	Code        string     `json:"code" db:"code"`
	Description string     `json:"description" db:"description"`
	Category    string     `json:"category" db:"category"`
	AppID       int        `json:"app_id" db:"app_id"`
	IsActive    bool       `json:"is_active" db:"is_active"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt   *time.Time `protobuf:"deleted_at" json:"deleted_at,omitempty" db:"deleted_at"`
}
