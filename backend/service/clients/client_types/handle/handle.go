package handle

import (
	"backend/protos/gen/go/clients/client_types"
	"backend/service/clients/client_types/models"
	"context"
	"google.golang.org/grpc"
	"log/slog"
)

type ClientTypesService interface {
	Create(ctx context.Context, req *models.CreateParams) (*models.ClientType, error)
	Get(ctx context.Context, id int) (*models.ClientType, error)
	Update(ctx context.Context, req *models.UpdateParams) (*models.ClientType, error)
	List(ctx context.Context, filter models.Filter, pagination models.Pagination) ([]*models.ClientType, int, error)
	Delete(ctx context.Context, id int, permanent bool) error
	Restore(ctx context.Context, id int) (*models.ClientType, error)
}

type serverAPI struct {
	client_types.UnimplementedClientTypeServiceServer
	service ClientTypesService
	logger  *slog.Logger
}

func Register(gRPCServer *grpc.Server, service ClientTypesService, logger *slog.Logger) {
	client_types.RegisterClientTypeServiceServer(gRPCServer, &serverAPI{
		service: service,
		logger:  logger,
	})
}
