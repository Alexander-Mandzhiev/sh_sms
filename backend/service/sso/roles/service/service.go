package service

import (
	"backend/service/sso/models"
	"context"
	"errors"
	"github.com/google/uuid"
	"log/slog"
)

var (
	ErrInvalidArgument  = errors.New("invalid argument")
	ErrInternal         = errors.New("internal server error")
	ErrPermissionDenied = errors.New("permission denied")
	ErrAlreadyExists    = errors.New("role already exists")
	ErrNotFound         = errors.New("role not found")
	ErrConflict         = errors.New("conflict fields")
)

type RolesProvider interface {
	Create(ctx context.Context, role *models.Role) error
	Update(ctx context.Context, role *models.Role) (*models.Role, error)

	GetByID(ctx context.Context, clientID, roleID uuid.UUID) (*models.Role, error)
	List(ctx context.Context, req models.ListRequest) ([]models.Role, int, error)

	HardDelete(ctx context.Context, clientID, roleID uuid.UUID) error
	SoftDelete(ctx context.Context, clientID, roleID uuid.UUID) error
	Restore(ctx context.Context, clientID, roleID uuid.UUID) (*models.Role, error)

	HasDependencies(ctx context.Context, roleID uuid.UUID) (bool, error)
	RoleExists(ctx context.Context, clientID uuid.UUID, name string) (bool, error)
}

type Service struct {
	logger   *slog.Logger
	provider RolesProvider
}

func New(provider RolesProvider, logger *slog.Logger) *Service {
	const op = "service.New"

	if logger == nil {
		logger = slog.Default()
	}

	logger.Info("initializing sso handle - service roles", slog.String("op", op))
	return &Service{
		provider: provider,
		logger:   logger,
	}
}
