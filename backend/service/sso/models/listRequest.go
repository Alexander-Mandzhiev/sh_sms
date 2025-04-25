package models

import (
	"github.com/google/uuid"
)

type ListRequest struct {
	Page           int
	Count          int
	ClientID       *uuid.UUID
	AppID          *int
	EmailFilter    *string
	PhoneFilter    *string
	NameFilter     *string
	LevelFilter    *int
	ActiveOnly     *bool
	CodeFilter     *string
	CategoryFilter *string
}
