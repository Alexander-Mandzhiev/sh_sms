package app_manager_handle

import (
	pb "backend/protos/gen/go/apps/app_manager"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (s *serverAPI) Create(ctx context.Context, req *pb.CreateRequest) (*pb.App, error) {
	const op = "handler.Create"
	logger := s.logger.With(slog.String("op", op))
	logger.Debug("Request received", slog.Any("request", req))

	if req.Code == "" || req.Name == "" {
		logger.Error("Empty required fields", slog.String("code", req.Code), slog.String("name", req.Name))
		return nil, status.Error(codes.InvalidArgument, "code and name are required")
	}

	app, err := s.service.Create(ctx, req)
	if err != nil {
		logger.Error("Create failed", slog.String("error", err.Error()))
		return nil, convertError(err)
	}

	logger.Info("Application created", slog.Int("id", int(app.Id)))
	return app, nil
}
