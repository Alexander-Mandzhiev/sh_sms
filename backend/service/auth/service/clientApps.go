package service

import (
	"context"
	"log/slog"

	"backend/protos/gen/go/apps/clients_apps"
)

type ClientAppsService interface {
	GetClientApp(ctx context.Context, req *client_apps.IdentifierRequest) (*client_apps.ClientApp, error)
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

func (s *clientAppsService) GetClientApp(ctx context.Context, req *client_apps.IdentifierRequest) (*client_apps.ClientApp, error) {
	s.logger.Debug("getting client app binding", "client_id", req.ClientId, "app_id", req.AppId)
	return s.client.GetClientApp(ctx, req)
}
