package handle

import (
	"backend/protos/gen/go/sso/permissions"
	"backend/service/sso/models"
	"context"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"log/slog"
)

type PermissionsService interface {
	Create(ctx context.Context, perm models.Permission) (*models.Permission, error)
	Get(ctx context.Context, id uuid.UUID, appID int) (*models.Permission, error)
	Update(ctx context.Context, perm models.Permission) (*models.Permission, error)
	Delete(ctx context.Context, id uuid.UUID, appID int, permanent bool) error
	List(ctx context.Context, filter models.ListRequest) ([]models.Permission, int, error)
	Restore(ctx context.Context, id uuid.UUID, appID int) (*models.Permission, error)
}

type serverAPI struct {
	permissions.UnimplementedPermissionServiceServer
	service PermissionsService
	logger  *slog.Logger
}

func Register(gRPCServer *grpc.Server, service PermissionsService, logger *slog.Logger) {
	permissions.RegisterPermissionServiceServer(gRPCServer, &serverAPI{
		service: service,
		logger:  logger,
	})
}
