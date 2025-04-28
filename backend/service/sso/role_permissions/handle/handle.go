package handle

import (
	"backend/protos/gen/go/sso/role_permissions"
	"backend/service/sso/models"
	"context"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"log/slog"
)

type RolePermissionService interface {
	AddPermissionsToRole(ctx context.Context, clientID, roleID uuid.UUID, appID int, permissionIDs []uuid.UUID) (*models.OperationStatus, error)
	RemovePermissionsFromRole(ctx context.Context, clientID, roleID uuid.UUID, appID int, permissionIDs []uuid.UUID) (*models.OperationStatus, error)
	ListPermissionsForRole(ctx context.Context, clientID, roleID uuid.UUID, appID int) ([]uuid.UUID, error)
	ListRolesForPermission(ctx context.Context, clientID, permissionID uuid.UUID, appID int) ([]uuid.UUID, error)
	HasPermission(ctx context.Context, clientID, roleID, permissionID uuid.UUID, appID int) (bool, error)
}

type serverAPI struct {
	role_permissions.UnimplementedRolePermissionServiceServer
	service RolePermissionService
	logger  *slog.Logger
}

func Register(gRPCServer *grpc.Server, service RolePermissionService, logger *slog.Logger) {
	role_permissions.RegisterRolePermissionServiceServer(gRPCServer, &serverAPI{
		service: service,
		logger:  logger,
	})
}
func convertUUIDsToStrings(ids []uuid.UUID) []string {
	result := make([]string, 0, len(ids))
	for _, id := range ids {
		result = append(result, id.String())
	}
	return result
}
