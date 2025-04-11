package service

import (
	pb "backend/protos/gen/go/apps/app_manager"
	"context"
	"log/slog"
)

func (s *Service) Get(ctx context.Context, req *pb.GetRequest) (*pb.App, error) {
	const op = "service.Get"
	logger := s.logger.With(slog.String("op", op))

	if req.GetId() == 0 && req.GetCode() == "" {
		logger.Error("No identifier provided")
		return nil, ErrIdentifierRequired
	}

	if req.GetId() != 0 && req.GetCode() != "" {
		logger.Error("Conflicting identifiers")
		return nil, ErrConflictParams
	}

	if req.GetId() < 0 {
		logger.Error("Invalid ID")
		return nil, ErrInvalidID
	}

	return s.provider.Get(ctx, req)
}
