package classes_handle

import (
	"backend/pkg/models/library"
	library "backend/protos/gen/go/library"
	"context"
	"google.golang.org/grpc"
	"log/slog"
)

type ClassesService interface {
	Get(ctx context.Context, id int) (*library_models.Class, error)
	List(ctx context.Context) ([]*library_models.Class, error)
}

type serverAPI struct {
	library.UnimplementedClassServiceServer
	service ClassesService
	logger  *slog.Logger
}

func Register(gRPCServer *grpc.Server, service ClassesService, logger *slog.Logger) {
	library.RegisterClassServiceServer(gRPCServer, &serverAPI{
		service: service,
		logger:  logger.With("module", "classes"),
	})
}
