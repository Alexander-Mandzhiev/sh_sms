package service

import (
	"backend/service/apps/models"
	"context"
	"log/slog"
)

type AppsProvider interface {
	Create(ctx context.Context, app *models.CreateApp) (*models.App, error)
	GetByID(ctx context.Context, id int) (*models.App, error)
	GetByCode(ctx context.Context, code string) (*models.App, error)
	Update(ctx context.Context, app *models.App) (*models.App, error)
	DeleteByID(ctx context.Context, id int) error
	DeleteByCode(ctx context.Context, code string) error
	List(ctx context.Context, filter models.ListFilter) ([]models.App, int, error)
}

type Service struct {
	provider AppsProvider
	logger   *slog.Logger
}

func New(provider AppsProvider, logger *slog.Logger) *Service {
	const op = "service.New"

	if logger == nil {
		logger = slog.Default()
	}

	logger.Info("initializing apps service", slog.String("op", op))
	return &Service{
		provider: provider,
		logger:   logger,
	}
}
