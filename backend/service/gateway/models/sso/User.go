package sso_models

import (
	"backend/protos/gen/go/sso/users"
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID  `json:"id"`
	ClientID  uuid.UUID  `json:"client_id"`
	Email     string     `json:"email"`
	FullName  string     `json:"full_name"`
	Phone     string     `json:"phone"`
	IsActive  bool       `json:"is_active"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

func UserFromProto(u *users.User) (*User, error) {
	id, err := uuid.Parse(u.Id)
	if err != nil {
		return nil, err
	}

	clientID, err := uuid.Parse(u.ClientId)
	if err != nil {
		return nil, err
	}

	var updatedAt, deletedAt *time.Time

	if u.UpdatedAt != nil {
		t := u.UpdatedAt.AsTime()
		updatedAt = &t
	}

	if u.DeletedAt != nil {
		t := u.DeletedAt.AsTime()
		deletedAt = &t
	}

	return &User{
		ID:        id,
		ClientID:  clientID,
		Email:     u.Email,
		FullName:  u.FullName,
		Phone:     u.Phone,
		IsActive:  u.IsActive,
		CreatedAt: u.CreatedAt.AsTime(),
		UpdatedAt: updatedAt,
		DeletedAt: deletedAt,
	}, nil
}
