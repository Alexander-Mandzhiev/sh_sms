package service

import (
	"context"
	"log/slog"

	"backend/protos/gen/go/sso/users_roles"
)

type UserRoleService interface {
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

func (s *userRoleService) ListForUser(ctx context.Context, req *user_roles.UserRequest) (*user_roles.UserRolesResponse, error) {
	s.logger.Debug("listing roles for user", "user_id", req.UserId, "client_id", req.ClientId, "app_id", req.AppId, "page", req.Page, "count", req.Count)
	return s.client.ListForUser(ctx, req)
}

func (s *userRoleService) ListForRole(ctx context.Context, req *user_roles.RoleRequest) (*user_roles.UserRolesResponse, error) {
	s.logger.Debug("listing users for role", "role_id", req.RoleId, "client_id", req.ClientId, "app_id", req.AppId, "page", req.Page, "count", req.Count)
	return s.client.ListForRole(ctx, req)
}
