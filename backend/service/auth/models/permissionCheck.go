package models

import (
	"backend/protos/gen/go/auth"
	"fmt"
	"github.com/google/uuid"
)

type PermissionCheck struct {
	UserID     uuid.UUID
	ClientID   uuid.UUID
	AppID      int
	Permission string
	Resource   string
}

func PermissionCheckFromProto(pb *auth.PermissionCheckRequest) (*PermissionCheck, error) {
	userID, err := uuid.Parse(pb.UserId)
	if err != nil {
		return nil, fmt.Errorf("invalid user_id: %w", err)
	}

	clientID, err := uuid.Parse(pb.ClientId)
	if err != nil {
		return nil, fmt.Errorf("invalid client_id: %w", err)
	}

	return &PermissionCheck{
		UserID:     userID,
		ClientID:   clientID,
		AppID:      int(pb.AppId),
		Permission: pb.Permission,
		Resource:   pb.Resource,
	}, nil
}
