package service

import (
	"context"
	"log/slog"

	"backend/protos/gen/go/sso/roles"
)

type RoleService interface {
	GetRole(ctx context.Context, req *roles.GetRequest) (*roles.Role, error)
	ListRoles(ctx context.Context, req *roles.ListRequest) (*roles.ListResponse, error)
}

type roleService struct {
	client roles.RoleServiceClient
	logger *slog.Logger
}

func NewRoleService(client roles.RoleServiceClient, logger *slog.Logger) RoleService {
	return &roleService{
		client: client,
		logger: logger.With("service", "role"),
	}
}

func (s *roleService) GetRole(ctx context.Context, req *roles.GetRequest) (*roles.Role, error) {
	s.logger.Debug("getting role", "role_id", req.Id, "client_id", req.ClientId, "app_id", req.AppId)
	return s.client.GetRole(ctx, req)
}

func (s *roleService) ListRoles(ctx context.Context, req *roles.ListRequest) (*roles.ListResponse, error) {
	s.logger.Debug("listing roles", "client_id", req.ClientId, "app_id", req.AppId, "filter_active", req.ActiveOnly)
	return s.client.ListRoles(ctx, req)
}
