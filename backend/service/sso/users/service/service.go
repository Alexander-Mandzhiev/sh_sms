package service

import (
	"backend/service/sso/models"
	"context"
	"errors"
	"github.com/google/uuid"
	"log/slog"
)

var (
	ErrInvalidArgument    = errors.New("invalid argument")
	ErrInternal           = errors.New("internal server error")
	ErrPermissionDenied   = errors.New("permission denied")
	ErrNotFound           = errors.New("user not found")
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrConflict           = errors.New("conflict fields")
)

type UsersProvider interface {
	Create(ctx context.Context, user *models.User) error
	Update(ctx context.Context, user *models.User) error
	UpdatePasswordHash(ctx context.Context, userID uuid.UUID, passwordHash string) error

	GetByID(ctx context.Context, clientID, userID uuid.UUID) (*models.User, error)
	GetByEmail(ctx context.Context, clientID uuid.UUID, email string) (*models.User, error)
	List(ctx context.Context, filter models.ListRequest) ([]models.User, int, error)

	SoftDeleteUser(ctx context.Context, clientID, userID uuid.UUID) error
	HardDeleteUser(ctx context.Context, clientID, userID uuid.UUID) error
	Exists(ctx context.Context, clientID, userID uuid.UUID) (bool, error)
	Restore(ctx context.Context, clientID, userID uuid.UUID) (*models.User, error)
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

	logger.Info("initializing sso service - handle user", slog.String("op", op))
	return &Service{
		provider: provider,
		logger:   logger,
	}
}
