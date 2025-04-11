package service

import (
	pb "backend/protos/gen/go/apps/clients_apps"
	"context"
	"errors"
	"log/slog"
)

var (
	ErrNotFound         = errors.New("not found")
	ErrInvalidArgument  = errors.New("invalid argument")
	ErrAlreadyExists    = errors.New("already exists")
	ErrInternal         = errors.New("internal server error")
	ErrPermissionDenied = errors.New("permission denied")
)

type ClientAppsProvider interface {
	Create(ctx context.Context, clientApp *pb.ClientApp) error
	Get(ctx context.Context, clientID string, appID int32) (*pb.ClientApp, error)
	Update(ctx context.Context, clientApp *pb.ClientApp) error
	Delete(ctx context.Context, clientID string, appID int32) error
	List(ctx context.Context, filter *pb.ListRequest_Filter, limit, offset int32) ([]*pb.ClientApp, int32, error)
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
