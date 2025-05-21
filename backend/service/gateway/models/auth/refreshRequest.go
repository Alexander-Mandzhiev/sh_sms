package auth_models

import "github.com/google/uuid"

type RefreshRequest struct {
	ClientID uuid.UUID `json:"client_id" validate:"required,uuid4"`
	AppID    int       `json:"app_id" validate:"required,min=1"`
}
