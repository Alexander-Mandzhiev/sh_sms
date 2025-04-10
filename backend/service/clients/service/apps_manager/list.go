package apps_manager_service

import (
	pb "backend/protos/gen/go/apps/app_manager"
	"context"
	"log/slog"
)

func (s *Service) List(ctx context.Context, req *pb.ListRequest) (*pb.ListResponse, error) {
	const op = "service.List"
	logger := s.logger.With(slog.String("op", op))

	_ = logger

	return nil, nil
}
