package service

import (
	"backend/service/apps/models"
	"context"
	"log/slog"
)

type ClientAppsProvider interface {
	Create(ctx context.Context, params models.CreateClientApp) (*models.ClientApp, error)
	Get(ctx context.Context, clientID string, appID int) (*models.ClientApp, error)
	Update(ctx context.Context, params models.UpdateClientApp) (*models.ClientApp, error)
	Delete(ctx context.Context, clientID string, appID int) error
	List(ctx context.Context, filter models.ListFilter) ([]*models.ClientApp, int, error)
}

type Service struct {
	logger   *slog.Logger
	provider ClientAppsProvider
}

func New(provider ClientAppsProvider, logger *slog.Logger) *Service {
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
