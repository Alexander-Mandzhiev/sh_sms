package service

import (
	pb "backend/protos/gen/go/apps/app_manager"
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
)

func (s *Service) Create(ctx context.Context, req *pb.CreateRequest) (*pb.App, error) {
	const op = "service.Create"
	logger := s.logger.With(slog.String("op", op))

	if req.GetCode() == "" || req.GetName() == "" {
		logger.Error("Empty code or name")
		return nil, ErrInvalidCredentials
	}

	now := timestamppb.Now()
	app := &pb.App{
		Code:        req.GetCode(),
		Name:        req.GetName(),
		Description: req.GetDescription(),
		IsActive:    req.GetIsActive(),
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	createdApp, err := s.provider.Create(ctx, app)
	if err != nil {
		logger.Error("Create failed", slog.Any("error", err))
		return nil, err
	}

	logger.Info("App created", slog.Int("id", int(createdApp.GetId())))
	return createdApp, nil
}
