package handle

import (
	"backend/protos/gen/go/sso/roles"
	"backend/service/sso/models"
	"context"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"log/slog"
)

type RoleService interface {
	Create(ctx context.Context, role *models.Role) error
	Get(ctx context.Context, clientID, roleID uuid.UUID, appID int) (*models.Role, error)
	Update(ctx context.Context, role *models.Role) (*models.Role, error)
	Delete(ctx context.Context, clientID, roleID uuid.UUID, appID int, permanent bool) error
	List(ctx context.Context, req models.ListRequest) ([]models.Role, int, error)
	Restore(ctx context.Context, clientID, roleID uuid.UUID, appID int) (*models.Role, error)
}

type serverAPI struct {
	roles.UnimplementedRoleServiceServer
	service RoleService
	logger  *slog.Logger
}

func Register(gRPCServer *grpc.Server, service RoleService, logger *slog.Logger) {
	roles.RegisterRoleServiceServer(gRPCServer, &serverAPI{
		service: service,
		logger:  logger,
	})
}
