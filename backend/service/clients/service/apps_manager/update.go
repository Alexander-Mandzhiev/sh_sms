package apps_manager_service

import (
	pb "backend/protos/gen/go/apps/app_manager"
	"context"
	"log/slog"
)

func (s *Service) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.App, error) {
	const op = "service.Update"
	logger := s.logger.With(slog.String("op", op))
	_ = logger
	return nil, nil
}
