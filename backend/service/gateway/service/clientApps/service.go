package client_app_service

import (
	"context"
	"log/slog"

	"backend/protos/gen/go/apps/clients_apps"
)

type ClientAppsService interface {
	CreateClientApp(ctx context.Context, req *client_apps.CreateRequest) (*client_apps.ClientApp, error)
	GetClientApp(ctx context.Context, req *client_apps.IdentifierRequest) (*client_apps.ClientApp, error)
	UpdateClientApp(ctx context.Context, req *client_apps.UpdateRequest) (*client_apps.ClientApp, error)
	DeleteClientApp(ctx context.Context, req *client_apps.IdentifierRequest) (*client_apps.DeleteResponse, error)
	ListClientApps(ctx context.Context, req *client_apps.ListRequest) (*client_apps.ListResponse, error)
}

type clientAppsService struct {
	client client_apps.ClientsAppServiceClient
	logger *slog.Logger
}

func NewClientAppsService(client client_apps.ClientsAppServiceClient, logger *slog.Logger) ClientAppsService {
	return &clientAppsService{
		client: client,
		logger: logger.With("service", "client_apps"),
	}
}

func (s *clientAppsService) CreateClientApp(ctx context.Context, req *client_apps.CreateRequest) (*client_apps.ClientApp, error) {
	s.logger.Debug("creating client app binding", "client_id", req.ClientId, "app_id", req.AppId, "is_active", req.GetIsActive())
	return s.client.CreateClientApp(ctx, req)
}

func (s *clientAppsService) GetClientApp(ctx context.Context, req *client_apps.IdentifierRequest) (*client_apps.ClientApp, error) {
	s.logger.Debug("getting client app binding", "client_id", req.ClientId, "app_id", req.AppId)
	return s.client.GetClientApp(ctx, req)
}

func (s *clientAppsService) UpdateClientApp(ctx context.Context, req *client_apps.UpdateRequest) (*client_apps.ClientApp, error) {
	s.logger.Debug("updating client app binding", "client_id", req.ClientId, "app_id", req.AppId, "is_active", req.GetIsActive())
	return s.client.UpdateClientApp(ctx, req)
}

func (s *clientAppsService) DeleteClientApp(ctx context.Context, req *client_apps.IdentifierRequest) (*client_apps.DeleteResponse, error) {
	s.logger.Debug("deleting client app binding", "client_id", req.ClientId, "app_id", req.AppId)
	return s.client.DeleteClientApp(ctx, req)
}

func (s *clientAppsService) ListClientApps(ctx context.Context, req *client_apps.ListRequest) (*client_apps.ListResponse, error) {
	s.logger.Debug("listing client app bindings", "filter_client_id", req.ClientId, "filter_app_id", req.AppId, "filter_is_active", req.IsActive, "page", req.Page, "count", req.Count)
	return s.client.ListClientsApp(ctx, req)
}
