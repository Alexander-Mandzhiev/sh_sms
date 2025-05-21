package sso_models

import "github.com/google/uuid"

type CreateRoleRequest struct {
	ClientID    string    `json:"client_id" binding:"required,uuid"`
	AppID       int       `json:"app_id" binding:"required,min=1"`
	Name        string    `json:"name" binding:"required,min=3,max=150"`
	Description string    `json:"description,omitempty" binding:"max=1000"`
	Level       int       `json:"level" binding:"required,min=1,max=1000"`
	IsCustom    *bool     `json:"is_custom,omitempty"`
	CreatedBy   uuid.UUID `json:"created_by" binding:"required,uuid"`
}
