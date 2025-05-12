package handle

import (
	pb "backend/protos/gen/go/apps/clients_apps"
	"backend/service/apps/models"
	"backend/service/utils"
	"context"
	"log/slog"
)

func (s *serverAPI) CreateClientApp(ctx context.Context, req *pb.CreateRequest) (*pb.ClientApp, error) {
	const op = "grpc.handler.ClientApp.Create"
	logger := s.logger.With(slog.String("op", op))
	logger.Debug("Create request received", slog.String("client_id", req.GetClientId()), slog.Int("app_id", int(req.GetAppId())))

	if err := utils.ValidateUUIDToString(req.GetClientId()); err != nil {
		logger.Warn("client_id validation failed", slog.Any("error", err))
		return nil, s.convertError(err)
	}
	if err := utils.ValidateAppID(int(req.GetAppId())); err != nil {
		logger.Warn("app_id validation failed", slog.Any("error", err))
		return nil, s.convertError(err)
	}

	params := models.CreateClientApp{
		ClientID: req.GetClientId(),
		AppID:    int(req.GetAppId()),
		IsActive: true,
	}

	if req.IsActive != nil {
		params.IsActive = *req.IsActive
	}

	createdApp, err := s.service.Create(ctx, params)
	if err != nil {
		return nil, s.convertError(err)
	}

	logger.Info("Client app created successfully", slog.String("client_id", createdApp.ClientID), slog.Int("app_id", int(createdApp.AppID)))
	return s.convertToPbClientApp(createdApp), nil
}
