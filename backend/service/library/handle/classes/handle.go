package classes_handle

import (
	"backend/protos/gen/go/library/classes"
	"context"
	"google.golang.org/grpc"
	"log/slog"
)

type ClassesService interface {
	Create(ctx context.Context, req *classes.CreateClassRequest) (*classes.Class, error)
	Get(ctx context.Context, req *classes.GetClassRequest) (*classes.Class, error)
	Update(ctx context.Context, req *classes.UpdateClassRequest) (*classes.Class, error)
	Delete(ctx context.Context, req *classes.DeleteClassRequest) error
	List(ctx context.Context, req *classes.ListClassesRequest) (*classes.ListClassesResponse, error)
}

type serverAPI struct {
	classes.UnimplementedClassesServiceServer
	service ClassesService
	logger  *slog.Logger
}

func Register(gRPCServer *grpc.Server, service ClassesService, logger *slog.Logger) {
	classes.RegisterClassesServiceServer(gRPCServer, &serverAPI{
		service: service,
		logger:  logger.With("module", "classes"),
	})
}
