package auth_models

import (
	"github.com/google/uuid"
)

type AllSessionsFilter struct {
	ClientID   uuid.UUID `json:"client_id"`
	AppID      int       `json:"app_id"`
	Page       int       `json:"page"`
	Count      int       `json:"count"`
	ActiveOnly *bool     `json:"active_only"`
	FullName   *string   `json:"full_name"`
	Phone      *string   `json:"phone"`
	Email      *string   `json:"email"`
}
