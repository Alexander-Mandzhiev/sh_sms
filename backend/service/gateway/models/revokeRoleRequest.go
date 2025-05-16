package models

type RevokeRoleRequest struct {
	UserID   string `json:"user_id" binding:"required,uuid"`
	RoleID   string `json:"role_id" binding:"required,uuid"`
	ClientID string `json:"client_id" binding:"required,uuid"`
	AppID    int32  `json:"app_id" binding:"required,min=1"`
}
