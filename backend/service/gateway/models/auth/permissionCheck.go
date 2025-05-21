package auth_models

import (
	"backend/pkg/utils"
	"backend/protos/gen/go/auth"
	"fmt"
	"github.com/google/uuid"
)

type PermissionCheck struct {
	UserID     uuid.UUID `json:"user_id"`
	ClientID   uuid.UUID `json:"client_id"`
	AppID      int       `json:"app_id"`
	Permission string    `json:"permission"`
	Resource   string    `json:"resource"`
}

func PermissionCheckFromProto(pb *auth.PermissionCheckRequest) (*PermissionCheck, error) {
	clientID, err := utils.ValidateAndReturnUUID(pb.GetClientId())
	if err != nil {
		return nil, err
	}
	userID, err := utils.ValidateAndReturnUUID(pb.GetUserId())
	if err != nil {
		return nil, err
	}

	if err = utils.ValidateAppID(int(pb.GetAppId())); err != nil {
		return nil, fmt.Errorf("invalid app_id: %w", err)
	}

	if pb.Permission == "" {
		return nil, fmt.Errorf("permission is required")
	}

	return &PermissionCheck{
		UserID:     userID,
		ClientID:   clientID,
		AppID:      int(pb.AppId),
		Permission: pb.Permission,
		Resource:   pb.Resource,
	}, nil
}
