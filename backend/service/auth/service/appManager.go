package service

import (
	"context"
	"log/slog"

	"backend/protos/gen/go/apps/app_manager"
)

type AppService interface {
	GetApp(ctx context.Context, req *app_manager.AppIdentifier) (*app_manager.App, error)
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
