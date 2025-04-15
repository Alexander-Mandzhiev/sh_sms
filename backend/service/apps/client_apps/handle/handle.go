package handle

import (
	pb "backend/protos/gen/go/apps/clients_apps"
	"backend/service/apps/models"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
)

type ClientAppService interface {
	Create(ctx context.Context, params models.CreateClientApp) (*models.ClientApp, error)
	Get(ctx context.Context, clientID string, appID int32) (*models.ClientApp, error)
	Update(ctx context.Context, clientID string, appID int32, isActive bool) (*models.ClientApp, error)
	Delete(ctx context.Context, clientID string, appID int32) error
	List(ctx context.Context, filter models.ListFilter) ([]models.ClientApp, int32, error)
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

func (s *serverAPI) convertToPbClientApp(app *models.ClientApp) *pb.ClientApp {
	return &pb.ClientApp{
		ClientId:  app.ClientID,
		AppId:     app.AppID,
		IsActive:  app.IsActive,
		CreatedAt: timestamppb.New(app.CreatedAt),
		UpdatedAt: timestamppb.New(app.UpdatedAt),
	}
}
