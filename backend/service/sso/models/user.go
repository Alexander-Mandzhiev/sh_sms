package models

import (
	"backend/protos/gen/go/sso/users"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type User struct {
	ID           uuid.UUID
	ClientID     uuid.UUID
	Email        string
	PasswordHash string
	FullName     string
	Phone        string
	IsActive     bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}

func ConvertUserToProto(u *User) *users.User {
	var deletedAt *timestamppb.Timestamp
	if u.DeletedAt != nil {
		deletedAt = timestamppb.New(*u.DeletedAt)
	}

	return &users.User{
		Id:        u.ID.String(),
		ClientId:  u.ClientID.String(),
		Email:     u.Email,
		FullName:  u.FullName,
		Phone:     u.Phone,
		IsActive:  u.IsActive,
		CreatedAt: timestamppb.New(u.CreatedAt),
		UpdatedAt: timestamppb.New(u.UpdatedAt),
		DeletedAt: deletedAt,
	}
}
