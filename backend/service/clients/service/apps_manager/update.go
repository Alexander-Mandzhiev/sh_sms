package apps_manager_service

import (
	"backend/protos/gen/go/apps/app_manager"
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
)

func (s *Service) Update(ctx context.Context, req *app_manager.UpdateRequest) (*app_manager.App, error) {
	const op = "service.Update"
	logger := s.logger.With(slog.String("op", op))

	if req.Name != nil {
		existingApp.Name = *req.Name
	}
	if req.Description != nil {
		existingApp.Description = *req.Description
	}
	if req.IsActive != nil {
		existingApp.IsActive = *req.IsActive
	}

	existingApp.UpdatedAt = timestamppb.Now()

	if err = s.provider.Update(ctx, existingApp); err != nil {
		logger.Error("failed to update app", slog.Any("error", err))
		return nil, err
	}

	return existingApp, nil
}
