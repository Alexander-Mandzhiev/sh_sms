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
}

func SessionFilterFromProto(pb *auth.SessionFilter) (*SessionFilter, error) {
	var userID uuid.UUID
	if pb.UserId != "" {
		var err error
		if userID, err = uuid.Parse(pb.UserId); err != nil {
			return nil, fmt.Errorf("invalid user_id: %w", err)
		}
	} else {
		userID = uuid.Nil
	}

	clientID, err := uuid.Parse(pb.ClientId)
	if err != nil {
		return nil, fmt.Errorf("invalid client_id: %w", err)
	}

	return &SessionFilter{
		UserID:     userID,
		ClientID:   clientID,
		AppID:      int(pb.AppId),
		ActiveOnly: pb.ActiveOnly,
	}, nil
}
