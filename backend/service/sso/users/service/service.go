package service

import (
	"backend/service/sso/models"
	"context"
	"github.com/google/uuid"
	"log/slog"
)

type UsersProvider interface {
	Create(ctx context.Context, user *models.User) error
	Get(ctx context.Context, clientID, userID uuid.UUID) (*models.User, error)
	GetByEmail(ctx context.Context, clientID uuid.UUID, email string) (*models.User, error)
	Update(ctx context.Context, userID, clientID uuid.UUID, update models.UserUpdate) error
	UpdatePasswordHash(ctx context.Context, userID uuid.UUID, passwordHash string) error
	SoftDeleteUser(ctx context.Context, clientID, userID uuid.UUID) error
	HardDeleteUser(ctx context.Context, clientID, userID uuid.UUID) error
	List(ctx context.Context, filter models.ListRequest) ([]models.User, int, error)
}

type Service struct {
	logger   *slog.Logger
	provider UsersProvider
}

func New(provider UsersProvider, logger *slog.Logger) *Service {
	const op = "service.New"

	if logger == nil {
		logger = slog.Default()
	}

	logger.Info("initializing client apps service", slog.String("op", op))
	return &Service{
		provider: provider,
		logger:   logger,
	}
}
