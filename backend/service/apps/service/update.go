package service

import (
	"backend/protos/gen/go/apps"
	"backend/service/apps/repository"
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (s *Service) Update(ctx context.Context, req *apps.UpdateRequest) (*apps.App, error) {
	const op = "service.Update"
	logger := s.logger.With(slog.String("op", op), slog.String("app_id", req.GetAppId()))
	if req.GetAppId() == "" {
		logger.Error("empty app_id")
		return nil, status.Error(codes.InvalidArgument, "app_id is required")
	}

	app := &apps.App{Id: req.GetAppId()}

	if req.Name != nil {
		name := req.GetName().GetValue()
		if name == "" {
			logger.Error("empty name not allowed")
			return nil, status.Error(codes.InvalidArgument, "name cannot be empty")
		}
		if len(name) > 255 {
			logger.Error("name too long", slog.Int("length", len(name)))
			return nil, status.Error(codes.InvalidArgument, "name must be <= 255 chars")
		}
		app.Name = name
	}

	if req.Description != nil {
		app.Description = req.GetDescription().GetValue()
	}

	if req.IsActive != nil {
		app.IsActive = req.GetIsActive().GetValue()
	}

	if req.Metadata != nil {
		app.Metadata = req.GetMetadata()
	}

	updatedApp, err := s.provider.Update(ctx, app)
	if err != nil {
		if errors.Is(err, repository.ErrAppNotFound) {
			logger.Error("app not found")
			return nil, status.Error(codes.NotFound, "app not found")
		}
		logger.Error("update failed", slog.String("error", err.Error()))
		return nil, status.Errorf(codes.Internal, "update failed")
	}

	logger.Info("app updated")
	return updatedApp, nil
}
