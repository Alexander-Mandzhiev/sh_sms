package models

import (
	"backend/protos/gen/go/auth"
	"fmt"
	"github.com/google/uuid"
)

type SessionFilter struct {
	UserID     uuid.UUID
	ClientID   uuid.UUID
	AppID      int
	ActiveOnly bool
	Page       int
	Count      int
}

func SessionFilterFromProto(pb *auth.SessionFilter) (*SessionFilter, error) {
	var userID uuid.UUID
	if pb.UserId != "" {
		var err error
		userID, err = uuid.Parse(pb.UserId)
		if err != nil {
			return nil, fmt.Errorf("invalid user_id: %w", err)
		}
	} else {
		userID = uuid.Nil
	}

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

	return &SessionFilter{
		UserID:     userID,
		ClientID:   clientID,
		AppID:      int(pb.AppId),
		ActiveOnly: pb.ActiveOnly,
		Page:       int(pb.Page),
		Count:      int(pb.Count),
	}, nil
}
