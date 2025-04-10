package apps_manager_service

import (
	"backend/protos/gen/go/apps/app_manager"
	"context"
	"log/slog"
)

func (s *Service) List(ctx context.Context, req *app_manager.ListRequest) (*app_manager.ListResponse, error) {
	const op = "service.List"
	logger := s.logger.With(slog.String("op", op))

	apps, total, err := s.provider.List(ctx, req)
	if err != nil {
		logger.Error("Failed to list apps",
			slog.Any("request", req),
			slog.Any("error", err))
		return nil, err
	}

	return &app_manager.ListResponse{
		Apps:       apps,
		TotalCount: total,
		Page:       req.Page,
		Count:      int32(len(apps)),
	}, nil
}
