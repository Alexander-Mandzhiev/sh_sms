package handle

import (
	sl "backend/pkg/logger"
	pb "backend/protos/gen/go/apps/clients_apps"
	"context"
	"log/slog"
)

func (s *serverAPI) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.ClientApp, error) {
	const op = "handler.Update"
	logger := s.logger.With(slog.String("op", op))
	logger.Debug("starting operation", slog.Any("request", req))

	resp, err := s.service.Update(ctx, req)
	if err != nil {
		logger.Error("operation failed", sl.Err(err, true))
		return nil, convertError(err)
	}

	logger.Info("operation completed successfully")
	return resp, nil
}
