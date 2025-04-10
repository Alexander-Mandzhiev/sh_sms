package apps_manager_service

import (
	pb "backend/protos/gen/go/apps/app_manager"
	"context"
	"log/slog"
)

type ClientAppsProvider interface {
	Create(ctx context.Context, clientApp *pb.CreateRequest) error
	Get(ctx context.Context, clientID string, appID int32) (*pb.App, error)
	Update(ctx context.Context, clientApp *pb.UpdateRequest) error
	Delete(ctx context.Context, clientID string, appID int32) error
	List(ctx context.Context, filter *pb.ListRequest, limit, offset int32) (*pb.ListResponse, int32, error)
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

	logger.Info("initializing apps service", slog.String("op", op))
	return &Service{
		provider: provider,
		logger:   logger,
	}
}
