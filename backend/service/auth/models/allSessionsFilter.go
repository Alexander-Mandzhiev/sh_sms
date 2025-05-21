package models

import (
	"backend/protos/gen/go/auth"
	"fmt"
	"github.com/google/uuid"
)

type AllSessionsFilter struct {
	ClientID   uuid.UUID
	AppID      int
	Page       int
	Count      int
	ActiveOnly *bool
	FullName   *string
	Phone      *string
	Email      *string
}

func AllSessionsFilterFromProto(pb *auth.AllSessionsFilter) (*AllSessionsFilter, error) {
	var clientID uuid.UUID
	if pb.ClientId != "" {
		var err error
		clientID, err = uuid.Parse(pb.ClientId)
		if err != nil {
			return nil, fmt.Errorf("invalid client_id: %w", err)
		}
	} else {
		clientID = uuid.Nil
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
