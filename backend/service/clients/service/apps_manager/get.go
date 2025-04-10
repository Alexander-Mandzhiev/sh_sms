package apps_manager_service

import (
	pb "backend/protos/gen/go/apps/app_manager"
	"context"
	"log/slog"
)

func (s *Service) Get(ctx context.Context, req *pb.GetRequest) (*pb.App, error) {
	const op = "service.Get"
	logger := s.logger.With(slog.String("op", op))

	_ = logger
	return nil, nil
}
