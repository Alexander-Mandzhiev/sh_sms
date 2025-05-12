package models

import "github.com/google/uuid"

type CreateUserRequest struct {
	ClientID uuid.UUID `json:"client_id" binding:"required"`
	Email    string    `json:"email" binding:"required,email"`
	Password string    `json:"password" binding:"required,min=8"`
	FullName string    `json:"full_name"`
	Phone    string    `json:"phone"`
}
