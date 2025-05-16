package models

import (
	"backend/protos/gen/go/auth"
)

type UserInfo struct {
	ID          string
	Email       string
	Phone       string
	FullName    string
	IsActive    bool
	Roles       []string
	Permissions []string
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
