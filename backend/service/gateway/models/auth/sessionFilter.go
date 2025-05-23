package auth_models

import (
	"github.com/google/uuid"
)

type SessionFilter struct {
	UserID     uuid.UUID `json:"user_id"`
	ClientID   uuid.UUID `json:"client_id"`
	AppID      int       `json:"app_id"`
	ActiveOnly bool      `json:"active_only"`
	Page       int       `json:"page"`
	Count      int       `json:"count"`
}
