package apps_manager_service

import (
	"backend/protos/gen/go/apps/app_manager"
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
)

func (s *Service) Create(ctx context.Context, req *app_manager.CreateRequest) (*app_manager.App, error) {
	const op = "service.Create"
	logger := s.logger.With(slog.String("op", op))

	now := timestamppb.Now()
	app := &app_manager.App{
		Code:        req.Code,
		Name:        req.Name,
		Description: req.Description,
		IsActive:    req.IsActive != nil && *req.IsActive,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	if err := s.provider.Create(ctx, app); err != nil {
		logger.Error("failed to create app", slog.Any("error", err))
		return nil, err
	}

	logger.Info("app created", slog.Int("id", int(app.Id)))
	return app, nil
}
