package auth_models

import (
	"backend/protos/gen/go/auth"
)

type UserInfo struct {
	ID          string   `json:"id"`
	Email       string   `json:"email"`
	FullName    string   `json:"full_name"`
	IsActive    bool     `json:"is_active"`
	Roles       []string `json:"roles"`
	Permissions []string `json:"permissions"`
}

func ConvertUserToProto(u *UserInfo) *auth.UserInfo {
	return &auth.UserInfo{
		Id:          u.ID,
		Email:       u.Email,
		FullName:    u.FullName,
		IsActive:    u.IsActive,
		Roles:       u.Roles,
		Permissions: u.Permissions,
	}
}
