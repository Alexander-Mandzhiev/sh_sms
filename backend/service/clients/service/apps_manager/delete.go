package apps_manager_service

import (
	"backend/protos/gen/go/apps/app_manager"
	"context"
	"log/slog"
)

func (s *Service) Delete(ctx context.Context, req *app_manager.DeleteRequest) (*app_manager.DeleteResponse, error) {
	const op = "service.Delete"
	logger := s.logger.With(slog.String("op", op))

	if err := s.provider.Delete(ctx, req.GetId()); err != nil {
		logger.Error("failed to delete app", slog.Any("error", err))
		return nil, err
	}

	return &app_manager.DeleteResponse{Success: true}, nil
}
