package service

import (
	"context"
	"log/slog"

	"backend/protos/gen/go/sso/users_roles"
)

type UserRoleService interface {
	AssignRole(ctx context.Context, req *user_roles.AssignRequest) (*user_roles.UserRole, error)
	RevokeRole(ctx context.Context, req *user_roles.RevokeRequest) (*user_roles.RevokeResponse, error)
	ListForUser(ctx context.Context, req *user_roles.UserRequest) (*user_roles.UserRolesResponse, error)
	ListForRole(ctx context.Context, req *user_roles.RoleRequest) (*user_roles.UserRolesResponse, error)
}

type userRoleService struct {
	client user_roles.UserRoleServiceClient
	logger *slog.Logger
}

func NewUserRoleService(client user_roles.UserRoleServiceClient, logger *slog.Logger) UserRoleService {
	return &userRoleService{
		client: client,
		logger: logger.With("service", "user_role"),
	}
}

func (s *userRoleService) AssignRole(ctx context.Context, req *user_roles.AssignRequest) (*user_roles.UserRole, error) {
	s.logger.Debug("assigning role to user", "user_id", req.UserId, "role_id", req.RoleId, "client_id", req.ClientId, "app_id", req.AppId, "expires_at", req.ExpiresAt)
	return s.client.Assign(ctx, req)
}

func (s *userRoleService) RevokeRole(ctx context.Context, req *user_roles.RevokeRequest) (*user_roles.RevokeResponse, error) {
	s.logger.Debug("revoking role from user", "user_id", req.UserId, "role_id", req.RoleId, "client_id", req.ClientId, "app_id", req.AppId)
	return s.client.Revoke(ctx, req)
}

func (s *userRoleService) ListForUser(ctx context.Context, req *user_roles.UserRequest) (*user_roles.UserRolesResponse, error) {
	s.logger.Debug("listing roles for user", "user_id", req.UserId, "client_id", req.ClientId, "app_id", req.AppId, "page", req.Page, "count", req.Count)
	return s.client.ListForUser(ctx, req)
}

func (s *userRoleService) ListForRole(ctx context.Context, req *user_roles.RoleRequest) (*user_roles.UserRolesResponse, error) {
	s.logger.Debug("listing users for role", "role_id", req.RoleId, "client_id", req.ClientId, "app_id", req.AppId, "page", req.Page, "count", req.Count)
	return s.client.ListForRole(ctx, req)
}
