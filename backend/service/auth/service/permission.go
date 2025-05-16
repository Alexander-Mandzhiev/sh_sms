package service

import (
	"context"
	"log/slog"

	"backend/protos/gen/go/sso/permissions"
)

type PermissionService interface {
	CreatePermission(ctx context.Context, req *permissions.CreateRequest) (*permissions.Permission, error)
	GetPermission(ctx context.Context, req *permissions.GetRequest) (*permissions.Permission, error)
	UpdatePermission(ctx context.Context, req *permissions.UpdateRequest) (*permissions.Permission, error)
	DeletePermission(ctx context.Context, req *permissions.DeleteRequest) (*permissions.SuccessResponse, error)
	ListPermissions(ctx context.Context, req *permissions.ListRequest) (*permissions.ListResponse, error)
	RestorePermission(ctx context.Context, req *permissions.RestoreRequest) (*permissions.Permission, error)
}

type permissionService struct {
	client permissions.PermissionServiceClient
	logger *slog.Logger
}

func NewPermissionService(client permissions.PermissionServiceClient, logger *slog.Logger) PermissionService {
	return &permissionService{
		client: client,
		logger: logger.With("service", "permission"),
	}
}

func (s *permissionService) CreatePermission(ctx context.Context, req *permissions.CreateRequest) (*permissions.Permission, error) {
	s.logger.Debug("creating permission", "code", req.Code, "category", req.Category, "app_id", req.AppId)
	return s.client.CreatePermission(ctx, req)
}

func (s *permissionService) GetPermission(ctx context.Context, req *permissions.GetRequest) (*permissions.Permission, error) {
	s.logger.Debug("getting permission", "permission_id", req.Id, "app_id", req.AppId)
	return s.client.GetPermission(ctx, req)
}

func (s *permissionService) UpdatePermission(ctx context.Context, req *permissions.UpdateRequest) (*permissions.Permission, error) {
	s.logger.Debug("updating permission", "permission_id", req.Id, "app_id", req.AppId)
	return s.client.UpdatePermission(ctx, req)
}

func (s *permissionService) DeletePermission(ctx context.Context, req *permissions.DeleteRequest) (*permissions.SuccessResponse, error) {
	s.logger.Debug("deleting permission", "permission_id", req.Id, "app_id", req.AppId, "permanent", req.Permanent)
	return s.client.DeletePermission(ctx, req)
}

func (s *permissionService) ListPermissions(ctx context.Context, req *permissions.ListRequest) (*permissions.ListResponse, error) {
	s.logger.Debug("listing permissions", "app_id", req.AppId, "category", req.Category, "active_only", req.ActiveOnly)
	return s.client.ListPermissions(ctx, req)
}

func (s *permissionService) RestorePermission(ctx context.Context, req *permissions.RestoreRequest) (*permissions.Permission, error) {
	s.logger.Debug("restoring permission", "permission_id", req.Id, "app_id", req.AppId)
	return s.client.RestorePermission(ctx, req)
}
