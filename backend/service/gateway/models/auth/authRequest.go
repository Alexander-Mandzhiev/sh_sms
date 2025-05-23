package auth_models

import (
	"github.com/google/uuid"
)

type AuthRequest struct {
	ClientID uuid.UUID `json:"client_id"`
	AppID    int       `json:"app_id"`
	Login    string    `json:"login"`
	Password string    `json:"password"`
}
