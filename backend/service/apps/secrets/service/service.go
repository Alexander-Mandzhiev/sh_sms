package service

import (
	"backend/service/apps/models"
	"context"
	"log/slog"
	"time"
)

type SecretsProvider interface {
	Create(ctx context.Context, params models.CreateSecretParams) (*models.Secret, error)
	Get(ctx context.Context, clientID string, appID int, secretType string) (*models.Secret, error)
	Rotate(ctx context.Context, params models.RotateSecretParams) (*models.Secret, error)
	Revoke(ctx context.Context, clientID string, appID int, secretType string) (*models.Secret, error)
	Delete(ctx context.Context, clientID string, appID int, secretType string) error
	List(ctx context.Context, filter models.ListFilter) ([]*models.Secret, int, error)
	GetRotation(ctx context.Context, clientID string, appID int, secretType string, rotatedAt time.Time) (*models.RotationHistory, error)
	ListRotations(ctx context.Context, filter models.ListFilter) ([]*models.RotationHistory, int, error)
}

type Service struct {
	logger   *slog.Logger
	provider SecretsProvider
}

func New(provider SecretsProvider, logger *slog.Logger) *Service {
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
