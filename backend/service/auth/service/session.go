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
	GetSession(ctx context.Context, userID uuid.UUID, clientID uuid.UUID, appID int, refreshTokenHash string) (*models.Session, error)
	RevokeSession(ctx context.Context, sessionID uuid.UUID) error
	ListSessions(ctx context.Context, filter models.SessionFilter) ([]models.Session, error)
	DeleteSession(ctx context.Context, sessionID uuid.UUID) error
	GetSessionByTokenHash(ctx context.Context, accessTokenHash, refreshTokenHash string) (*models.Session, error)
	GetSessionByToken(ctx context.Context, token string) (*models.Session, error)
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
