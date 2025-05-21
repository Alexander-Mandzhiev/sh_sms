package sso_models

type UpdateRoleRequest struct {
	ClientID    string  `json:"client_id"`
	AppID       int32   `json:"app_id"`
	Name        *string `json:"name,omitempty" binding:"omitempty,min=3,max=64"`
	Description *string `json:"description,omitempty" binding:"omitempty,max=255"`
	Level       *int32  `json:"level,omitempty" binding:"omitempty,min=1,max=1000"`
}
