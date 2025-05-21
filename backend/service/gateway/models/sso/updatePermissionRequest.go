package sso_models

type UpdatePermissionRequest struct {
	Code        *string `json:"code,omitempty" binding:"omitempty,alphanum,max=50"`
	Description *string `json:"description,omitempty" binding:"omitempty,max=255"`
	Category    *string `json:"category,omitempty" binding:"omitempty,max=50"`
	IsActive    *bool   `json:"is_active,omitempty"`
}
