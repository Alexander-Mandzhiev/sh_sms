package sso_models

type SetPasswordRequest struct {
	NewPassword string `json:"new_password" binding:"required,min=8,max=72"`
}
