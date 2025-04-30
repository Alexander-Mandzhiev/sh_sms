package handle

import (
	user_roles "backend/protos/gen/go/sso/users_roles"
	"backend/service/sso/models"
	"context"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"log/slog"
)

type UserRoleService interface {
	Assign(ctx context.Context, role *models.UserRole) (*models.UserRole, error)
	Revoke(ctx context.Context, userID uuid.UUID, roleID uuid.UUID, clientID uuid.UUID, appID int) error
	ListForUser(ctx context.Context, filter models.ListRequest) ([]*models.UserRole, int, error)
	ListForRole(ctx context.Context, filter models.ListRequest) ([]*models.UserRole, int, error)
}

type serverAPI struct {
	user_roles.UnimplementedUserRoleServiceServer
	service UserRoleService
	logger  *slog.Logger
}

func Register(gRPCServer *grpc.Server, service UserRoleService, logger *slog.Logger) {
	user_roles.RegisterUserRoleServiceServer(gRPCServer, &serverAPI{
		service: service,
		logger:  logger,
	})
}
