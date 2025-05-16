package models

import (
	"backend/protos/gen/go/sso/users"
	"github.com/google/uuid"
)

type UserInfo struct {
	ID       uuid.UUID
	Email    string
	FullName string
	Phone    string
	IsActive bool
}

func ConvertUserInfoToProto(u *UserInfo) *users.UserInfo {
	return &users.UserInfo{
		Id:       u.ID.String(),
		Email:    u.Email,
		FullName: u.FullName,
		Phone:    u.Phone,
		IsActive: u.IsActive,
	}
}
