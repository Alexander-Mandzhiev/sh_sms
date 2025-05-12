package role_permission_service

import (
	"context"
	"log/slog"

	"backend/protos/gen/go/sso/role_permissions"
)

type RolePermissionService interface {
	AddPermissionsToRole(ctx context.Context, req *role_permissions.PermissionsRequest) (*role_permissions.OperationStatus, error)
	RemovePermissionsFromRole(ctx context.Context, req *role_permissions.PermissionsRequest) (*role_permissions.OperationStatus, error)
	ListPermissionsForRole(ctx context.Context, req *role_permissions.ListPermissionsRequest) (*role_permissions.ListPermissionsResponse, error)
	ListRolesForPermission(ctx context.Context, req *role_permissions.ListRolesRequest) (*role_permissions.ListRolesResponse, error)
	HasPermission(ctx context.Context, req *role_permissions.HasPermissionRequest) (*role_permissions.HasPermissionResponse, error)
}

type rolePermissionService struct {
	client role_permissions.RolePermissionServiceClient
	logger *slog.Logger
}

func NewRolePermissionService(client role_permissions.RolePermissionServiceClient, logger *slog.Logger) RolePermissionService {
	return &rolePermissionService{
		client: client,
		logger: logger.With("service", "role_permission"),
	}
}

func (s *rolePermissionService) AddPermissionsToRole(ctx context.Context, req *role_permissions.PermissionsRequest) (*role_permissions.OperationStatus, error) {
	s.logger.Debug("adding permissions to role", "role_id", req.RoleId, "client_id", req.ClientId, "app_id", req.AppId, "permissions_count", len(req.PermissionIds))
	return s.client.AddPermissionsToRole(ctx, req)
}

func (s *rolePermissionService) RemovePermissionsFromRole(ctx context.Context, req *role_permissions.PermissionsRequest) (*role_permissions.OperationStatus, error) {
	s.logger.Debug("removing permissions from role", "role_id", req.RoleId, "client_id", req.ClientId, "app_id", req.AppId, "permissions_count", len(req.PermissionIds))
	return s.client.RemovePermissionsFromRole(ctx, req)
}

func (s *rolePermissionService) ListPermissionsForRole(ctx context.Context, req *role_permissions.ListPermissionsRequest) (*role_permissions.ListPermissionsResponse, error) {
	s.logger.Debug("listing permissions for role", "role_id", req.RoleId, "client_id", req.ClientId, "app_id", req.AppId)
	return s.client.ListPermissionsForRole(ctx, req)
}

func (s *rolePermissionService) ListRolesForPermission(ctx context.Context, req *role_permissions.ListRolesRequest) (*role_permissions.ListRolesResponse, error) {
	s.logger.Debug("listing roles for permission", "permission_id", req.PermissionId, "client_id", req.ClientId, "app_id", req.AppId)
	return s.client.ListRolesForPermission(ctx, req)
}

func (s *rolePermissionService) HasPermission(ctx context.Context, req *role_permissions.HasPermissionRequest) (*role_permissions.HasPermissionResponse, error) {
	s.logger.Debug("checking permission", "role_id", req.RoleId, "permission_id", req.PermissionId, "client_id", req.ClientId, "app_id", req.AppId)
	return s.client.HasPermission(ctx, req)
}
