package handle

import (
	sl "backend/pkg/logger"
	pb "backend/protos/gen/go/apps/app_manager"
	"backend/service/apps/models"
	"context"
	"errors"
	"log/slog"
	"time"
)

func (s *serverAPI) UpdateApp(ctx context.Context, req *pb.UpdateRequest) (*pb.App, error) {
	const op = "grpc.handler.AppManager.Update"
	id := req.GetId()
	logger := s.logger.With(slog.String("op", op), slog.Int("id", int(req.Id)), slog.Time("timestamp", time.Now()))

	if err := validateUpdateRequest(req); err != nil {
		logger.Warn("Validation failed", sl.Err(err, true))
		return nil, s.convertError(err)
	}

	updateParams := models.UpdateApp{
		Code:        req.Code,
		Name:        req.Name,
		Description: req.Description,
		IsActive:    req.IsActive,
	}

	app, err := s.service.Update(ctx, int(id), updateParams)
	if err != nil {
		logger.Error("Failed to update app", sl.Err(err, true), slog.String("error_type", errors.Unwrap(err).Error()))
		return nil, s.convertError(err)
	}

	logger.Info("App updated successfully", slog.Int("app_id", app.ID))
	return s.convertAppToProto(app), nil
}
