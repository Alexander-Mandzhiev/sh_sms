package classes_service

import (
	"backend/pkg/models/library"
	"context"
	"log/slog"
)

type ClassesProvider interface {
	GetClassByID(ctx context.Context, id int) (*library_models.Class, error)
	ListClasses(ctx context.Context) ([]*library_models.Class, error)
}

type Service struct {
	logger   *slog.Logger
	provider ClassesProvider
}

func New(provider ClassesProvider, logger *slog.Logger) *Service {
	const op = "service.New.Library.Classes"

	if logger == nil {
		logger = slog.Default()
	}

	logger.Info("initializing library handle - service classes", slog.String("op", op))
	return &Service{
		provider: provider,
		logger:   logger,
	}
}
