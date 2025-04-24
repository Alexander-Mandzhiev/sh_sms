package handle

import (
	pb "backend/protos/gen/go/apps/clients_apps"
	"backend/service/utils"
	"context"
	"log/slog"
)

func (s *serverAPI) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.ClientApp, error) {
	const op = "grpc.handler.ClientApp.Update"
	logger := s.logger.With(slog.String("op", op), slog.String("client_id", req.GetClientId()), slog.Int("app_id", int(req.GetAppId())))
	logger.Debug("starting operation")

	if err := utils.ValidateUUIDToString(req.GetClientId()); err != nil {
		logger.Warn("client_id validation failed", slog.Any("error", err))
		return nil, s.convertError(err)
	}
	if err := utils.ValidateAppID(int(req.GetAppId())); err != nil {
		logger.Warn("app_id validation failed", slog.Any("error", err))
		return nil, s.convertError(err)
	}

	updatedApp, err := s.service.Update(ctx, req.GetClientId(), int(req.GetAppId()), req.IsActive)
	if err != nil {
		return nil, s.convertError(err)
	}

	pbClientApp := s.convertToPbClientApp(updatedApp)
	logger.Info("client app updated successfully", slog.Bool("is_active", updatedApp.IsActive), slog.Time("updated_at", updatedApp.UpdatedAt))
	return pbClientApp, nil
}
