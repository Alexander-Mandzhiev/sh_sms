package auth_models

import (
	"backend/pkg/utils"
	"backend/protos/gen/go/auth"
	"fmt"
	"github.com/google/uuid"
)

type SessionFilter struct {
	UserID     uuid.UUID `json:"user_id"`
	ClientID   uuid.UUID `json:"client_id"`
	AppID      int       `json:"app_id"`
	ActiveOnly bool      `json:"active_only"`
	Page       int       `json:"page"`
	Count      int       `json:"count"`
}

func SessionFilterFromProto(pb *auth.SessionFilter) (*SessionFilter, error) {
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

	if err = utils.ValidatePagination(int(pb.GetPage()), int(pb.GetCount())); err != nil {
		return nil, err
	}

	return &SessionFilter{
		UserID:     userID,
		ClientID:   clientID,
		AppID:      int(pb.AppId),
		ActiveOnly: pb.ActiveOnly,
		Page:       int(pb.Page),
		Count:      int(pb.Count),
	}, nil
}
