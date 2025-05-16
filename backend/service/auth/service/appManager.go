package service

import (
	"context"
	"log/slog"

	"backend/protos/gen/go/apps/app_manager"
)

type AppService interface {
	CreateApp(ctx context.Context, req *app_manager.CreateRequest) (*app_manager.App, error)
	GetApp(ctx context.Context, req *app_manager.AppIdentifier) (*app_manager.App, error)
	UpdateApp(ctx context.Context, req *app_manager.UpdateRequest) (*app_manager.App, error)
	DeleteApp(ctx context.Context, req *app_manager.AppIdentifier) (*app_manager.DeleteResponse, error)
	ListApps(ctx context.Context, req *app_manager.ListRequest) (*app_manager.ListResponse, error)
}

type appService struct {
	client app_manager.AppServiceClient
	logger *slog.Logger
}

func NewAppService(client app_manager.AppServiceClient, logger *slog.Logger) AppService {
	return &appService{
		client: client,
		logger: logger.With("service", "app_manager"),
	}
}

func (s *appService) CreateApp(ctx context.Context, req *app_manager.CreateRequest) (*app_manager.App, error) {
	s.logger.Debug("creating application", "code", req.Code, "name", req.Name, "is_active", req.IsActive)
	return s.client.CreateApp(ctx, req)
}

func (s *appService) GetApp(ctx context.Context, req *app_manager.AppIdentifier) (*app_manager.App, error) {
	logFields := []any{"identifier", "unknown"}

	switch id := req.Identifier.(type) {
	case *app_manager.AppIdentifier_Id:
		logFields = []any{"app_id", id.Id}
	case *app_manager.AppIdentifier_Code:
		logFields = []any{"app_code", id.Code}
	}

	s.logger.Debug("getting application", logFields...)
	return s.client.GetApp(ctx, req)
}

func (s *appService) UpdateApp(ctx context.Context, req *app_manager.UpdateRequest) (*app_manager.App, error) {
	s.logger.Debug("updating application", "app_id", req.Id, "code", req.Code, "name", req.Name, "is_active", req.IsActive)
	return s.client.UpdateApp(ctx, req)
}

func (s *appService) DeleteApp(ctx context.Context, req *app_manager.AppIdentifier) (*app_manager.DeleteResponse, error) {
	logFields := []any{"identifier", "unknown"}

	switch id := req.Identifier.(type) {
	case *app_manager.AppIdentifier_Id:
		logFields = []any{"app_id", id.Id}
	case *app_manager.AppIdentifier_Code:
		logFields = []any{"app_code", id.Code}
	}

	s.logger.Debug("deleting application", logFields...)
	return s.client.DeleteApp(ctx, req)
}

func (s *appService) ListApps(ctx context.Context, req *app_manager.ListRequest) (*app_manager.ListResponse, error) {
	s.logger.Debug("listing applications", "page", req.Page, "count", req.Count, "filter_active", req.FilterIsActive)
	return s.client.ListApps(ctx, req)
}
