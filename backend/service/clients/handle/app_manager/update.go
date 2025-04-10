package app_manager_handle

import (
	pb "backend/protos/gen/go/apps/app_manager"
	apps "backend/service/clients/service/apps_manager"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (s *serverAPI) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.App, error) {
	const op = "handler.Update"
	logger := s.logger.With(slog.String("op", op))
	logger.Debug("Request received", slog.Int("id", int(req.Id)))

	if req.Id <= 0 {
		logger.Error("Invalid ID")
		return nil, convertError(apps.ErrInvalidID)
	}

	if req.Name == nil && req.Description == nil && req.IsActive == nil {
		logger.Error("No fields to update")
		return nil, convertError(apps.ErrNoUpdateFields)
	}

	if req.Name != nil && *req.Name == "" {
		logger.Error("Empty name")
		return nil, status.Error(codes.InvalidArgument, "name cannot be empty")
	}

	app, err := s.service.Update(ctx, req)
	if err != nil {
		logger.Error("Update failed", slog.String("error", err.Error()))
		return nil, convertError(err)
	}

	logger.Info("Application updated", slog.Int("id", int(app.Id)))
	return app, nil
}
