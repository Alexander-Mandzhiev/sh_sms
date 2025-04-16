package handle

import (
	pb "backend/protos/gen/go/apps/clients_apps"
	"backend/service/apps/models"
	"context"
	"google.golang.org/grpc"
	"log/slog"
)

type ClientAppService interface {
	Create(ctx context.Context, params models.CreateClientApp) (*models.ClientApp, error)
	Get(ctx context.Context, clientID string, appID int) (*models.ClientApp, error)
	Update(ctx context.Context, clientID string, appID int, isActive *bool) (*models.ClientApp, error)
	Delete(ctx context.Context, clientID string, appID int) error
	List(ctx context.Context, filter models.ListFilter) ([]*models.ClientApp, int, error)
}

type serverAPI struct {
	pb.UnimplementedClientsAppServiceServer
	service ClientAppService
	logger  *slog.Logger
}

func Register(gRPCServer *grpc.Server, service ClientAppService, logger *slog.Logger) {
	pb.RegisterClientsAppServiceServer(gRPCServer, &serverAPI{
		service: service,
		logger:  logger,
	})
}
