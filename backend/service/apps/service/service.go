package service

import (
	"backend/protos/gen/go/apps"
	"context"
	"log/slog"
)

type AppsProvider interface {
	Create(ctx context.Context, app *apps.App) (*apps.App, error)
	Get(ctx context.Context, id string) (*apps.App, error)
	Update(ctx context.Context, app *apps.App) (*apps.App, error)
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, limit int32, offset int32, nameFilter string, activeOnly bool) ([]*apps.App, int32, error)
	GenerateSecretKey(ctx context.Context, appID string, generatedBy string, keyLength int32) (string, error)
	RotateSecretKey(ctx context.Context, appID string, rotatedBy string, invalidatePrevious bool) (string, error)
	RevokeSecretKey(ctx context.Context, appID string, revokedBy string, regenerate bool) (string, error)
	GetKeyRotationHistory(ctx context.Context, appID string, limit int32) ([]*apps.KeyRotationRecord, error)
}

type Service struct {
	logger   *slog.Logger
	provider AppsProvider
}

func New(appsProvider AppsProvider, logger *slog.Logger) *Service {
	const op = "service.New"
	if appsProvider == nil {
		logger.Error("apps provider is nil", slog.String("op", op))
		panic("apps provider cannot be nil")
	}
	logger.Info("service initialized", slog.String("op", op))
	return &Service{provider: appsProvider, logger: logger}
}
