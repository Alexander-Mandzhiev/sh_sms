package service

import (
	pb "backend/protos/gen/go/apps/app_manager"
	"context"
	"errors"
	"log/slog"
)

var (
	ErrInvalidID          = errors.New("invalid id")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrConflictParams     = errors.New("conflicting parameters")
	ErrInvalidPagination  = errors.New("invalid pagination parameters")
	ErrIdentifierRequired = errors.New("either id or code must be provided")
)

type AppsProvider interface {
	Create(ctx context.Context, app *pb.App) (*pb.App, error)
	Get(ctx context.Context, req *pb.GetRequest) (*pb.App, error)
	Update(ctx context.Context, app *pb.App) (*pb.App, error)
	Delete(ctx context.Context, id int32) error
	List(ctx context.Context, req *pb.ListRequest) (*pb.ListResponse, error)
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
