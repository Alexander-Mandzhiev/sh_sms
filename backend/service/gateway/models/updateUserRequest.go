package models

type UpdateUserRequest struct {
	ClientID string  `json:"client_id" binding:"required,uuid"`
	Email    *string `json:"email,omitempty" binding:"omitempty,email"`
	FullName *string `json:"full_name,omitempty"`
	Phone    *string `json:"phone,omitempty" binding:"omitempty,e164"`
}
