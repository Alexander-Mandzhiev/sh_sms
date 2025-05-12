package service

import (
	"context"
	"log/slog"

	"backend/protos/gen/go/sso/roles"
)

type RoleService interface {
	CreateRole(ctx context.Context, req *roles.CreateRequest) (*roles.Role, error)
	GetRole(ctx context.Context, req *roles.GetRequest) (*roles.Role, error)
	UpdateRole(ctx context.Context, req *roles.UpdateRequest) (*roles.Role, error)
	DeleteRole(ctx context.Context, req *roles.DeleteRequest) (*roles.DeleteResponse, error)
	ListRoles(ctx context.Context, req *roles.ListRequest) (*roles.ListResponse, error)
	RestoreRole(ctx context.Context, req *roles.RestoreRequest) (*roles.Role, error)
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

func (s *roleService) CreateRole(ctx context.Context, req *roles.CreateRequest) (*roles.Role, error) {
	s.logger.Debug("creating role", "client_id", req.ClientId, "app_id", req.AppId, "name", req.Name)
	return s.client.CreateRole(ctx, req)
}

func (s *roleService) GetRole(ctx context.Context, req *roles.GetRequest) (*roles.Role, error) {
	s.logger.Debug("getting role", "role_id", req.Id, "client_id", req.ClientId, "app_id", req.AppId)
	return s.client.GetRole(ctx, req)
}

func (s *roleService) UpdateRole(ctx context.Context, req *roles.UpdateRequest) (*roles.Role, error) {
	s.logger.Debug("updating role", "role_id", req.Id, "client_id", req.ClientId, "app_id", req.AppId)
	return s.client.UpdateRole(ctx, req)
}

func (s *roleService) DeleteRole(ctx context.Context, req *roles.DeleteRequest) (*roles.DeleteResponse, error) {
	s.logger.Debug("deleting role", "role_id", req.Id, "client_id", req.ClientId, "app_id", req.AppId, "permanent", req.Permanent)
	return s.client.DeleteRole(ctx, req)
}

func (s *roleService) ListRoles(ctx context.Context, req *roles.ListRequest) (*roles.ListResponse, error) {
	s.logger.Debug("listing roles", "client_id", req.ClientId, "app_id", req.AppId, "filter_active", req.ActiveOnly)
	return s.client.ListRoles(ctx, req)
}

func (s *roleService) RestoreRole(ctx context.Context, req *roles.RestoreRequest) (*roles.Role, error) {
	s.logger.Debug("restoring role", "role_id", req.Id, "client_id", req.ClientId, "app_id", req.AppId)
	return s.client.RestoreRole(ctx, req)
}
