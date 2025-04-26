package service

import (
	"backend/service/sso/models"
	"context"
	"errors"
	"github.com/google/uuid"
	"log/slog"
)

var (
	ErrAlreadyExists   = errors.New("permission already exists")
	ErrNotFound        = errors.New("permission not found")
	ErrInvalidArgument = errors.New("invalid argument")
	ErrInvalidState    = errors.New("invalid object state")
)

type PermissionProvider interface {
	Create(ctx context.Context, perm models.Permission) (*models.Permission, error)
	GetByID(ctx context.Context, id uuid.UUID, appID int) (*models.Permission, error)
	Update(ctx context.Context, perm models.Permission) (*models.Permission, error)
	List(ctx context.Context, filter models.ListRequest) ([]models.Permission, int, error)
	SoftDelete(ctx context.Context, id uuid.UUID, appID int) error
	HardDelete(ctx context.Context, id uuid.UUID, appID int) error
	Restore(ctx context.Context, id uuid.UUID, appID int) (*models.Permission, error)
	ExistByCode(ctx context.Context, code string, appID int) (bool, error)
	ExistByID(ctx context.Context, id uuid.UUID, appID int) (bool, error)
}

type Service struct {
	logger   *slog.Logger
	provider PermissionProvider
}

func New(provider PermissionProvider, logger *slog.Logger) *Service {
	const op = "service.New.Permission"

	if logger == nil {
		logger = slog.Default()
	}

	logger.Info("initializing sso handle - service permission", slog.String("op", op))
	return &Service{
		provider: provider,
		logger:   logger,
	}
}
