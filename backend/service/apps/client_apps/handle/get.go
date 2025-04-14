package handle

import (
	pb "backend/protos/gen/go/apps/clients_apps"
	"context"
	"log/slog"
)

func (s *serverAPI) Get(ctx context.Context, req *pb.GetRequest) (*pb.ClientApp, error) {
	const op = "handler.Get"
	logger := s.logger.With(slog.String("op", op))
	logger.Debug("starting operation", slog.Any("request", req))

	if err := validateClientID(req.ClientId); err != nil {
		return nil, err
	}
	if err := validateAppID(req.AppId); err != nil {
		return nil, err
	}

	resp, err := s.service.Get(ctx, req)
	if err != nil {
		return nil, s.handleError(op, err)
	}

	logger.Debug("operation completed", slog.Any("response", resp))
	return resp, nil
}
