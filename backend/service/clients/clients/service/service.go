package service

import (
	"backend/service/clients/clients/models"
	"context"
	"github.com/google/uuid"
	"log/slog"
)

type ClientsProvider interface {
	Create(ctx context.Context, client *models.CreateParams) (*models.Client, error)
	GetByID(ctx context.Context, id uuid.UUID) (*models.Client, error)
	List(ctx context.Context, filter models.Filter, pagination models.Pagination) ([]*models.Client, int, error)
	Update(ctx context.Context, params *models.UpdateParams) (*models.Client, error)
	Delete(ctx context.Context, id uuid.UUID, permanent bool) error
	Restore(ctx context.Context, id uuid.UUID) (*models.Client, error)
}

type Service struct {
	logger   *slog.Logger
	provider ClientsProvider
}

func New(provider ClientsProvider, logger *slog.Logger) *Service {
	const op = "service.New.Clients"
	if logger == nil {
		logger = slog.Default()
	}

	logger.Info("initializing clients service - handle clients", slog.String("op", op))
	return &Service{
		provider: provider,
		logger:   logger,
	}
}
