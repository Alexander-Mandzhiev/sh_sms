package client_apps_handle

import (
	sl "backend/pkg/logger"
	pb "backend/protos/gen/go/apps/clients_apps"
	"context"
	"log/slog"
)

func (s *serverAPI) Get(ctx context.Context, req *pb.GetRequest) (*pb.ClientApp, error) {
	const op = "handler.Get"
	logger := s.logger.With(slog.String("op", op))
	logger.Debug("starting operation", slog.Any("request", req))

	resp, err := s.service.Get(ctx, req)
	if err != nil {
		logger.Error("operation failed", sl.Err(err, true))
		return nil, convertError(err)
	}

	logger.Debug("operation completed", slog.Any("response", resp))
	return resp, nil
}
