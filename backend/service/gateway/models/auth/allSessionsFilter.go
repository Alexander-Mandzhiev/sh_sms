package auth_models

import (
	"backend/pkg/utils"
	"backend/protos/gen/go/auth"
	"fmt"
	"github.com/google/uuid"
)

type AllSessionsFilter struct {
	ClientID   uuid.UUID `json:"client_id"`
	AppID      int       `json:"app_id"`
	Page       int       `json:"page"`
	Count      int       `json:"count"`
	ActiveOnly *bool     `json:"active_only"`
	FullName   *string   `json:"full_name"`
	Phone      *string   `json:"phone"`
	Email      *string   `json:"email"`
}

func AllSessionsFilterFromProto(pb *auth.AllSessionsFilter) (*AllSessionsFilter, error) {
	clientID, err := utils.ValidateAndReturnUUID(pb.GetClientId())
	if err != nil {
		return nil, err
	}

	if err = utils.ValidateAppID(int(pb.GetAppId())); err != nil {
		return nil, fmt.Errorf("invalid app_id: %w", err)
	}

	if err = utils.ValidatePagination(int(pb.GetPage()), int(pb.GetCount())); err != nil {
		return nil, err
	}

	filter := &AllSessionsFilter{
		ClientID: clientID,
		AppID:    int(pb.AppId),
		Page:     int(pb.Page),
		Count:    int(pb.Count),
	}

	if pb.ActiveOnly != nil {
		filter.ActiveOnly = pb.ActiveOnly
	}
	if pb.FullName != nil {
		filter.FullName = pb.FullName
	}
	if pb.Phone != nil {
		filter.Phone = pb.Phone
	}
	if pb.Email != nil {
		filter.Email = pb.Email
	}

	return filter, nil
}
