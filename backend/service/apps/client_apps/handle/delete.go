package handle

import (
	pb "backend/protos/gen/go/apps/clients_apps"
	"backend/service/utils"
	"context"
	"log/slog"
)

func (s *serverAPI) DeleteClientApp(ctx context.Context, req *pb.IdentifierRequest) (*pb.DeleteResponse, error) {
	const op = "grpc.handler.ClientApp.Delete"
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

	if err := s.service.Delete(ctx, req.GetClientId(), int(req.GetAppId())); err != nil {
		return nil, s.convertError(err)
	}

	logger.Info("client app deleted successfully")
	return &pb.DeleteResponse{Success: true}, nil
}
