package handle

import (
	pb "backend/protos/gen/go/apps/app_manager"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (s *serverAPI) Get(ctx context.Context, req *pb.GetRequest) (*pb.App, error) {
	const op = "handler.Get"
	logger := s.logger.With(slog.String("op", op))
	logger.Debug("Get request received", slog.Any("request", req))

	if req.GetId() == 0 && req.GetCode() == "" {
		logger.Error("No identifier provided")
		return nil, status.Error(codes.InvalidArgument, "id or code must be provided")
	}

	app, err := s.service.Get(ctx, req)
	if err != nil {
		logger.Error("Get failed", slog.String("error", err.Error()))
		return nil, convertError(err)
	}

	logger.Debug("Application found", slog.Int("id", int(app.GetId())))
	return app, nil
}
