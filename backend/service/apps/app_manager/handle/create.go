package handle

import (
	sl "backend/pkg/logger"
	pb "backend/protos/gen/go/apps/app_manager"
	"backend/service/apps/models"
	"context"
	"log/slog"
	"time"
)

func (s *serverAPI) CreateApp(ctx context.Context, req *pb.CreateRequest) (*pb.App, error) {
	const op = "grpc.handler.AppManager.Create"
	logger := s.logger.With(slog.String("op", op), slog.Time("timestamp", time.Now()))

	if err := validateName(req.Name, 250); err != nil {
		logger.Warn("Name validation failed", sl.Err(err, true))
		return nil, s.convertError(err)
	}

	if err := validateCode(req.Code, 50); err != nil {
		logger.Warn("Code validation failed", sl.Err(err, true))
		return nil, s.convertError(err)
	}

	createParams := models.CreateApp{
		Code:        req.GetCode(),
		Name:        req.GetName(),
		Description: req.GetDescription(),
		IsActive:    true,
	}

	if req.IsActive != nil {
		createParams.IsActive = *req.IsActive
	}

	app, err := s.service.Create(ctx, &createParams)
	if err != nil {
		logger.Error("Create operation failed", sl.Err(err, true), slog.String("code", createParams.Code))
		return nil, s.convertError(err)
	}

	logger.Info("App created successfully", slog.Int("app_id", app.ID), slog.String("app_code", app.Code))
	return s.convertAppToProto(app), nil
}
