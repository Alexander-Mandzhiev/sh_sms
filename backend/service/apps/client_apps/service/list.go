package service

import (
	sl "backend/pkg/logger"
	pb "backend/protos/gen/go/apps/clients_apps"
	"context"
	"fmt"
	"log/slog"
)

func (s *Service) List(ctx context.Context, req *pb.ListRequest) (*pb.ListResponse, error) {
	const op = "service.List"
	logger := s.logger.With(slog.String("op", op))
	if req.Page < 1 || req.Count < 1 || req.Count > 100 {
		return nil, fmt.Errorf("%s: %w: invalid pagination params", op, ErrInvalidArgument)
	}

	offset := (req.Page - 1) * req.Count

	clientApps, total, err := s.provider.List(ctx, req.Filter, req.Count, offset)
	if err != nil {
		logger.Error("failed to list client apps", sl.Err(err, true))
		return nil, fmt.Errorf("%s: %w", op, ErrInternal)
	}

	return &pb.ListResponse{
		ClientApps: clientApps,
		TotalCount: total,
		Page:       req.Page,
		Count:      int32(len(clientApps)),
	}, nil
}
