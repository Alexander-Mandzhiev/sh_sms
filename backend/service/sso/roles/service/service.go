package service

import (
	"backend/service/sso/models"
	"context"
	"github.com/google/uuid"
	"log/slog"
)

type RolesProvider interface {
	Create(ctx context.Context, role *models.Role) error
	GetByID(ctx context.Context, clientID, roleID uuid.UUID) (*models.Role, error)
	HasDependencies(ctx context.Context, roleID uuid.UUID) (bool, error)
	HardDelete(ctx context.Context, clientID, roleID uuid.UUID) error
	Update(ctx context.Context, role *models.Role) (*models.Role, error)
	List(ctx context.Context, req models.ListRequest) ([]models.Role, int, error)
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
