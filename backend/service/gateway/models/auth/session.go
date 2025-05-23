package auth_models

import (
	"backend/protos/gen/go/auth"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net"
	"time"
)

type Session struct {
	SessionID    uuid.UUID  `json:"session_id"`
	UserID       uuid.UUID  `json:"user_id"`
	ClientID     uuid.UUID  `json:"client_id"`
	AppID        int        `json:"app_id"`
	IPAddress    net.IP     `json:"ip_address"`
	UserAgent    string     `json:"user_agent"`
	CreatedAt    time.Time  `json:"created_at"`
	LastActivity time.Time  `json:"last_activity"`
	ExpiresAt    time.Time  `json:"expires_at"`
	RevokedAt    *time.Time `json:"revoked_at"`
	FullName     string     `json:"full_name"`
	Phone        string     `json:"phone"`
	Email        string     `json:"email"`
}

func (s Session) IsActive() bool {
	return s.RevokedAt == nil && s.ExpiresAt.After(time.Now())
}

func SessionsFromProto(pbSessions []*auth.Session) []Session {
	var result []Session

	for _, pb := range pbSessions {
		if pb == nil {
			continue
		}

		session, err := ParseSessionFromProto(pb)
		if err != nil {
			continue
		}

		result = append(result, *session)
	}

	return result
}

func ParseSessionFromProto(pb *auth.Session) (*Session, error) {
	sessionID, err := uuid.Parse(pb.SessionId)
	if err != nil {
		return nil, fmt.Errorf("invalid session_id: %w", err)
	}

	userID, err := uuid.Parse(pb.UserId)
	if err != nil {
		return nil, fmt.Errorf("invalid user_id: %w", err)
	}

	clientID, err := uuid.Parse(pb.ClientId)
	if err != nil {
		return nil, fmt.Errorf("invalid client_id: %w", err)
	}

	return &Session{
		SessionID:    sessionID,
		UserID:       userID,
		ClientID:     clientID,
		AppID:        int(pb.AppId),
		IPAddress:    net.ParseIP(pb.IpAddress),
		UserAgent:    pb.UserAgent,
		CreatedAt:    pb.CreatedAt.AsTime(),
		LastActivity: pb.LastActivity.AsTime(),
		ExpiresAt:    pb.ExpiresAt.AsTime(),
		RevokedAt:    parseNullableTime(pb.RevokedAt),
		FullName:     pb.FullName,
		Phone:        pb.Phone,
		Email:        pb.Email,
	}, nil
}

func parseNullableTime(t *timestamppb.Timestamp) *time.Time {
	if t == nil {
		return nil
	}
	res := t.AsTime()
	return &res
}
