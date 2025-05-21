package sso_models

import "github.com/google/uuid"

type CreateUserRequest struct {
	ClientID uuid.UUID `json:"client_id" binding:"required,uuid"`
	Email    string    `json:"email" binding:"required,email"`
	Password string    `json:"password" binding:"required,min=8"`
	FullName string    `json:"full_name" binding:"required"`
	Phone    string    `json:"phone" binding:"required,e164"`
}
