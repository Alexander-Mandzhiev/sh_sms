package handle

import (
	pb "backend/protos/gen/go/apps/app_manager"
	"context"
	"google.golang.org/grpc"
	"log/slog"
)

type AppService interface {
	Create(ctx context.Context, req *pb.CreateRequest) (*pb.App, error)
	Get(ctx context.Context, req *pb.GetRequest) (*pb.App, error)
	Update(ctx context.Context, req *pb.UpdateRequest) (*pb.App, error)
	Delete(ctx context.Context, id int32) error
	List(ctx context.Context, req *pb.ListRequest) (*pb.ListResponse, error)
}

type serverAPI struct {
	pb.UnimplementedAppServiceServer
	service AppService
	logger  *slog.Logger
}

func Register(gRPCServer *grpc.Server, service AppService, logger *slog.Logger) {
	pb.RegisterAppServiceServer(gRPCServer, &serverAPI{
		service: service,
		logger:  logger,
	})
}
