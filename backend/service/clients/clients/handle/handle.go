package handle

import (
	"backend/protos/gen/go/clients/clients"
	"backend/service/clients/clients/models"
	"context"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"log/slog"
)

type ClientsService interface {
	Create(ctx context.Context, req *models.CreateParams) (*models.Client, error)
	Get(ctx context.Context, id uuid.UUID) (*models.Client, error)
	List(ctx context.Context, filter models.Filter, pagination models.Pagination) ([]*models.Client, int, error)
	Update(ctx context.Context, req *models.UpdateParams) (*models.Client, error)
	Delete(ctx context.Context, id uuid.UUID, permanent bool) error
	Restore(ctx context.Context, id uuid.UUID) (*models.Client, error)
}

type serverAPI struct {
	clients.UnimplementedClientServiceServer
	service ClientsService
	logger  *slog.Logger
}

func Register(gRPCServer *grpc.Server, service ClientsService, logger *slog.Logger) {
	clients.RegisterClientServiceServer(gRPCServer, &serverAPI{
		service: service,
		logger:  logger,
	})
}
