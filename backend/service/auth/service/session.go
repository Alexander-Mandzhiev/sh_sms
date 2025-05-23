package service

import (
	cfg "backend/pkg/config/auth"
	"backend/service/auth/models"
	"context"
	"errors"
	"github.com/google/uuid"
	"log/slog"
	"time"
)

var (
	ErrSessionNotFound     = errors.New("session not found")
	ErrTokenAlreadyRevoked = errors.New("token already revoked")
	ErrInvalidToken        = errors.New("invalid token")
)

type SessionProvider interface {
	CreateSession(ctx context.Context, session *models.Session) error
	UpdateTokens(ctx context.Context, sessionID uuid.UUID, accessHash, refreshHash string, expiresAt time.Time) error
	RevokeSession(ctx context.Context, sessionID uuid.UUID) error
	DeleteSession(ctx context.Context, sessionID uuid.UUID) error

	GetSession(ctx context.Context, userID uuid.UUID, clientID uuid.UUID, appID int, refreshTokenHash string) (*models.Session, error)
	GetSessionByToken(ctx context.Context, token string) (*models.Session, error)

	ListSessionsForUser(ctx context.Context, filter models.SessionFilter, fullName, phone, email string) ([]models.Session, error)
	ListAllSessions(ctx context.Context, filter models.AllSessionsFilter) ([]models.Session, error)
}

type SessionService struct {
	client SessionProvider
	logger *slog.Logger
	cfg    cfg.Config
}

func NewSessionService(client SessionProvider, logger *slog.Logger) *SessionService {
	return &SessionService{
		client: client,
		logger: logger.With("service", "sessions"),
	}
}
