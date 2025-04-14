package handle

import (
	pb "backend/protos/gen/go/apps/clients_apps"
	"context"
	"log/slog"
)

func (s *serverAPI) Create(ctx context.Context, req *pb.CreateRequest) (*pb.ClientApp, error) {
	const op = "handler.Create"
	logger := s.logger.With(slog.String("op", op))
	logger.Debug("starting operation", slog.Any("request", req))

	if err := validateClientID(req.ClientId); err != nil {
		return nil, err
	}
	if err := validateAppID(req.AppId); err != nil {
		return nil, err
	}

	resp, err := s.service.Create(ctx, req)
	if err != nil {
		return nil, s.handleError(op, err)
	}

	logger.Info("operation completed successfully")
	return resp, nil
}
