package handle

import (
	sl "backend/pkg/logger"
	pb "backend/protos/gen/go/apps/clients_apps"
	"context"
	"log/slog"
)

func (s *serverAPI) List(ctx context.Context, req *pb.ListRequest) (*pb.ListResponse, error) {
	const op = "handler.List"
	logger := s.logger.With(slog.String("op", op))
	logger.Debug("starting operation", slog.Any("request", req))

	resp, err := s.service.List(ctx, req)
	if err != nil {
		logger.Error("operation failed", sl.Err(err, true))
		return nil, convertError(err)
	}

	logger.Debug("operation completed", slog.Int("count", len(resp.ClientApps)))
	return resp, nil
}
