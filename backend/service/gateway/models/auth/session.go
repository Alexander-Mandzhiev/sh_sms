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

func SessionToProto(s *Session) (*auth.Session, error) {
	if s.SessionID == uuid.Nil {
		return nil, fmt.Errorf("invalid session_id")
	}

	return &auth.Session{
		SessionId:    s.SessionID.String(),
		CreatedAt:    timestamppb.New(s.CreatedAt),
		LastActivity: timestamppb.New(s.LastActivity),
		ClientId:     s.ClientID.String(),
		IpAddress:    s.IPAddress.String(),
		UserAgent:    s.UserAgent,
		UserId:       s.UserID.String(),
		FullName:     s.FullName,
		Phone:        s.Phone,
		Email:        s.Email,
	}, nil
}

func SessionsToProto(sessions []Session) (*auth.SessionList, error) {
	pbSessions := make([]*auth.Session, 0, len(sessions))
	for _, session := range sessions {
		protoSession, err := SessionToProto(&session)
		if err != nil {
			return nil, err
		}
		pbSessions = append(pbSessions, protoSession)
	}
	return &auth.SessionList{Sessions: pbSessions}, nil
}
