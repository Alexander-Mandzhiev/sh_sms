package models

import (
	"backend/protos/gen/go/auth"
	"fmt"
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
	FullName         string
	Phone            string
	Email            string
}

func (s Session) IsActive() bool {
	return s.RevokedAt == nil && s.ExpiresAt.After(time.Now())
}

func SessionToProto(s *Session) (*auth.Session, error) {
	if s.SessionID == uuid.Nil {
		return nil, fmt.Errorf("invalid session_id")
	}

	return &auth.Session{
		SessionId:    s.SessionID.String(),
		CreatedAt:    timestamppb.New(s.CreatedAt),
		LastActivity: timestamppb.New(s.LastActivity),
		ClientId:     s.ClientID.String(),
		AppId:        int32(s.AppID),
		IpAddress:    s.IPAddress.String(),
		UserAgent:    s.UserAgent,
		UserId:       s.UserID.String(),
		FullName:     s.FullName,
		Phone:        s.Phone,
		Email:        s.Email,
		ExpiresAt:    timestamppb.New(s.ExpiresAt),
		RevokedAt:    parseNullableTime(s.RevokedAt),
	}, nil
}

func parseNullableTime(t *time.Time) *timestamppb.Timestamp {
	if t == nil {
		return nil
	}
	return timestamppb.New(*t)
}

func SessionsToProto(sessions []Session) (*auth.SessionList, error) {
	pbSessions := make([]*auth.Session, 0, len(sessions))
	for _, session := range sessions {
		protoSession, err := SessionToProto(&session)
		if err != nil {
			return nil, fmt.Errorf("failed to convert session: %w", err)
		}
		pbSessions = append(pbSessions, protoSession)
	}
	return &auth.SessionList{Sessions: pbSessions}, nil
}
