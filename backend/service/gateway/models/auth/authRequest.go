package auth_models

import (
	"github.com/google/uuid"
)

type AuthRequest struct {
	ClientID uuid.UUID `json:"clientID"`
	AppID    int       `json:"appID"`
	Login    string    `json:"login"`
	Password string    `json:"password"`
}
