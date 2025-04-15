package handle

import (
	pb "backend/protos/gen/go/apps/app_manager"
	"backend/service/apps/models"
	"context"
	"google.golang.org/grpc"
	"log/slog"
)

type AppService interface {
	Create(ctx context.Context, params *models.CreateApp) (*models.App, error)
	GetByID(ctx context.Context, id int) (*models.App, error)
	GetByCode(ctx context.Context, code string) (*models.App, error)
	Update(ctx context.Context, id int, params models.UpdateApp) (*models.App, error)
	DeleteByID(ctx context.Context, id int) error
	DeleteByCode(ctx context.Context, code string) error
	List(ctx context.Context, filter models.ListFilter) ([]models.App, int, error)
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
