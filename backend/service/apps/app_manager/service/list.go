package service

import (
	pb "backend/protos/gen/go/apps/app_manager"
	"context"
	"log/slog"
)

func (s *Service) List(ctx context.Context, req *pb.ListRequest) (*pb.ListResponse, error) {
	const op = "service.List"
	logger := s.logger.With(slog.String("op", op))

	if req.GetPage() <= 0 || req.GetCount() <= 0 {
		logger.Error("Invalid page or count")
		return nil, ErrInvalidPagination
	}

	return s.provider.List(ctx, req)
}
