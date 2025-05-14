// models/permission.go
package models

import (
	"time"

	"backend/protos/gen/go/sso/permissions"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Permission struct {
	ID          string
	Code        string
	Description string
	Category    string
	AppID       int32
	IsActive    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

func PermissionFromProto(pb *permissions.Permission) (*Permission, error) {
	var deletedAt *time.Time
	if pb.DeletedAt != nil {
		deletedAtVal := pb.DeletedAt.AsTime()
		deletedAt = &deletedAtVal
	}

	return &Permission{
		ID:          pb.Id,
		Code:        pb.Code,
		Description: pb.Description,
		Category:    pb.Category,
		AppID:       pb.AppId,
		IsActive:    pb.IsActive,
		CreatedAt:   pb.CreatedAt.AsTime(),
		UpdatedAt:   pb.UpdatedAt.AsTime(),
		DeletedAt:   deletedAt,
	}, nil
}

func PermissionToProto(p *Permission) *permissions.Permission {
	var deletedAt *timestamppb.Timestamp
	if p.DeletedAt != nil {
		deletedAt = timestamppb.New(*p.DeletedAt)
	}

	return &permissions.Permission{
		Id:          p.ID,
		Code:        p.Code,
		Description: p.Description,
		Category:    p.Category,
		AppId:       p.AppID,
		IsActive:    p.IsActive,
		CreatedAt:   timestamppb.New(p.CreatedAt),
		UpdatedAt:   timestamppb.New(p.UpdatedAt),
		DeletedAt:   deletedAt,
	}
}
