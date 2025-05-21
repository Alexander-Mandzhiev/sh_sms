package sso_models

import (
	"backend/protos/gen/go/sso/roles"
	"time"
)

type Role struct {
	ID          string     `json:"id"`
	ClientID    string     `json:"client_id"`
	AppID       int32      `json:"app_id"`
	Name        string     `json:"name"`
	Description *string    `json:"description,omitempty"`
	Level       int32      `json:"level"`
	IsActive    bool       `json:"is_active"`
	IsCustom    bool       `json:"is_custom"`
	CreatedBy   *string    `json:"created_by,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}

func RoleFromProto(pbRole *roles.Role) (*Role, error) {
	var deletedAt *time.Time
	if pbRole.DeletedAt != nil {
		deletedAt = new(time.Time)
		*deletedAt = pbRole.DeletedAt.AsTime()
	}

	return &Role{
		ID:          pbRole.Id,
		ClientID:    pbRole.ClientId,
		AppID:       pbRole.AppId,
		Name:        pbRole.Name,
		Description: pbRole.Description,
		Level:       pbRole.Level,
		IsActive:    pbRole.IsActive,
		IsCustom:    pbRole.IsCustom,
		CreatedBy:   pbRole.CreatedBy,
		CreatedAt:   pbRole.CreatedAt.AsTime(),
		UpdatedAt:   pbRole.UpdatedAt.AsTime(),
		DeletedAt:   deletedAt,
	}, nil
}
