package handle

import (
	pb "backend/protos/gen/go/apps/clients_apps"
	"context"
	"log/slog"
)

func (s *serverAPI) Get(ctx context.Context, req *pb.IdentifierRequest) (*pb.ClientApp, error) {
	const op = "grpc.handler.ClientApp.Get"
	logger := s.logger.With(slog.String("op", op), slog.String("client_id", req.GetClientId()), slog.Int("app_id", int(req.GetAppId())))
	logger.Debug("starting operation")

	if err := validateClientID(req.GetClientId()); err != nil {
		logger.Warn("validation failed", slog.Any("error", err))
		return nil, s.convertError(err)
	}
	if err := validateAppID(req.GetAppId()); err != nil {
		logger.Warn("validation failed", slog.Any("error", err))
		return nil, s.convertError(err)
	}

	clientApp, err := s.service.Get(ctx, req.GetClientId(), req.GetAppId())
	if err != nil {
		return nil, s.convertError(err)
	}

	pbClientApp := s.convertToPbClientApp(clientApp)
	logger.Info("operation completed successfully", slog.Bool("is_active", clientApp.IsActive), slog.Time("created_at", clientApp.CreatedAt))
	return pbClientApp, nil
}
