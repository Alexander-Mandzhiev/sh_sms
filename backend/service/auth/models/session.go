package models

import (
	"backend/protos/gen/go/auth"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net"
	"time"
)

type Session struct {
	SessionID        uuid.UUID
	UserID           uuid.UUID
	ClientID         uuid.UUID
	AppID            int
	AccessTokenHash  string
	RefreshTokenHash string
	IPAddress        net.IP
	UserAgent        string
	CreatedAt        time.Time
	LastActivity     time.Time
	ExpiresAt        time.Time
	RevokedAt        *time.Time
}

func (s Session) IsActive() bool {
	return s.RevokedAt == nil && s.ExpiresAt.After(time.Now())
}

func SessionToProto(s *Session) *auth.Session {
	return &auth.Session{
		SessionId:    s.SessionID.String(),
		CreatedAt:    timestamppb.New(s.CreatedAt),
		LastActivity: timestamppb.New(s.LastActivity),
		ClientId:     s.ClientID.String(),
		IpAddress:    s.IPAddress.String(),
		UserAgent:    s.UserAgent,
	}
}
func SessionsToProto(sessions []Session) *auth.SessionList {
	pbSessions := make([]*auth.Session, 0, len(sessions))
	for _, session := range sessions {
		pbSessions = append(pbSessions, SessionToProto(&session))
	}
	return &auth.SessionList{Sessions: pbSessions}
}
