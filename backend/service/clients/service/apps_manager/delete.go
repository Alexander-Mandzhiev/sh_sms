package apps_manager_service

import (
	pb "backend/protos/gen/go/apps/app_manager"
	"context"
	"log/slog"
)

func (s *Service) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	const op = "service.Delete"
	logger := s.logger.With(slog.String("op", op))
	_ = logger

	return &pb.DeleteResponse{Success: true}, nil
}
