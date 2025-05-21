package sso_models

import "time"

type AssignRoleRequest struct {
	UserID     string     `json:"user_id" binding:"required,uuid"`
	RoleID     string     `json:"role_id" binding:"required,uuid"`
	ClientID   string     `json:"client_id" binding:"required,uuid"`
	AppID      int32      `json:"app_id" binding:"required,min=1"`
	ExpiresAt  *time.Time `json:"expires_at,omitempty"`
	AssignedBy string     `json:"assigned_by" binding:"required"`
}
