package sso_models

import (
	"backend/protos/gen/go/sso/users_roles"
	"time"
)

type UserRole struct {
	UserID     string     `json:"user_id"`
	RoleID     string     `json:"role_id"`
	ClientID   string     `json:"client_id"`
	AppID      int32      `json:"app_id"`
	AssignedBy string     `json:"assigned_by"`
	ExpiresAt  *time.Time `json:"expires_at,omitempty"`
	AssignedAt time.Time  `json:"assigned_at"`
}

func UserRoleFromProto(pb *user_roles.UserRole) *UserRole {
	var expiresAt *time.Time
	if pb.ExpiresAt != nil {
		t := pb.ExpiresAt.AsTime()
		expiresAt = &t
	}

	return &UserRole{
		UserID:     pb.UserId,
		RoleID:     pb.RoleId,
		ClientID:   pb.ClientId,
		AppID:      pb.AppId,
		AssignedBy: pb.AssignedBy,
		ExpiresAt:  expiresAt,
		AssignedAt: pb.AssignedAt.AsTime(),
	}
}
