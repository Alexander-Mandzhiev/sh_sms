package client_apps_handle

import (
	pb "backend/protos/gen/go/apps/clients_apps"
	"context"
	"google.golang.org/grpc"
	"log/slog"
)

type ClientAppsService interface {
	Create(ctx context.Context, req *pb.CreateRequest) (*pb.ClientApp, error)
	Get(ctx context.Context, req *pb.GetRequest) (*pb.ClientApp, error)
	Update(ctx context.Context, req *pb.UpdateRequest) (*pb.ClientApp, error)
	Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error)
	List(ctx context.Context, req *pb.ListRequest) (*pb.ListResponse, error)
}

type serverAPI struct {
	pb.UnimplementedClientsAppServiceServer
	service ClientAppsService
	logger  *slog.Logger
}

func Register(gRPCServer *grpc.Server, service ClientAppsService, logger *slog.Logger) {
	pb.RegisterClientsAppServiceServer(gRPCServer, &serverAPI{service: service, logger: logger})
}
