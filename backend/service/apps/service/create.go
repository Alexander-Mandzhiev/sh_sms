package service

import (
	"backend/protos/gen/go/apps"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (s *Service) Create(ctx context.Context, req *apps.CreateRequest) (*apps.App, error) {
	const op = "service.Create"
	logger := s.logger.With(slog.String("op", op))

	if req.GetName() == "" || len(req.GetName()) > 255 || req.GetCreatedBy() == "" {
		logger.Error("invalid request", slog.Bool("name_empty", req.GetName() == ""), slog.Bool("name_too_long", len(req.GetName()) > 255), slog.Bool("created_by_empty", req.GetCreatedBy() == ""))
		return nil, status.Error(codes.InvalidArgument,
			"name (1-255 chars) and created_by are required")
	}

	app := &apps.App{
		Name:        req.GetName(),
		Description: req.GetDescription(),
		CreatedBy:   req.GetCreatedBy(),
		Metadata:    req.GetMetadata(),
		IsActive:    true,
	}

	createdApp, err := s.provider.Create(ctx, app)
	if err != nil {
		logger.Error("failed to create app", slog.String("error", err.Error()))
		return nil, status.Errorf(codes.Internal, "failed to create app")
	}

	logger.Info("app created", slog.String("app_id", createdApp.GetId()))
	return createdApp, nil
}
