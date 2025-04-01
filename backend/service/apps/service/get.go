package service

import (
	"backend/protos/gen/go/apps"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (s *Service) Get(ctx context.Context, req *apps.GetRequest) (*apps.App, error) {
	const op = "service.Get"
	logger := s.logger.With(slog.String("op", op), slog.String("app_id", req.GetAppId()))

	if req.GetAppId() == "" {
		logger.Error("app_id is required")
		return nil, status.Error(codes.InvalidArgument, "app_id is required")
	}

	app, err := s.provider.Get(ctx, req.GetAppId())
	if err != nil {
		logger.Error("failed to get app", slog.String("error", err.Error()))
		return nil, status.Errorf(codes.NotFound, "app not found")
	}

	logger.Debug("app retrieved")
	return app, nil
}
