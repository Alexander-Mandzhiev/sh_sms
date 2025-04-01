package service

import (
	"backend/protos/gen/go/apps"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (s *Service) List(ctx context.Context, req *apps.ListRequest) (*apps.ListResponse, error) {
	const op = "service.List"
	logger := s.logger.With(slog.String("op", op))
	if req.GetLimit() < 0 || req.GetOffset() < 0 {
		logger.Error("invalid pagination", slog.Int("limit", int(req.GetLimit())), slog.Int("offset", int(req.GetOffset())))
		return nil, status.Error(codes.InvalidArgument, "limit and offset must be >= 0")
	}

	limit := req.GetLimit()
	if limit > 1000 {
		logger.Warn("adjusting limit to 1000", slog.Int("original", int(limit)))
		limit = 1000
	}

	var (
		nameFilter string
		activeOnly = false
	)

	if req.GetNameFilter() != nil {
		nameFilter = req.GetNameFilter().GetValue()
	}

	if req.GetActiveOnly() != nil {
		activeOnly = req.GetActiveOnly().GetValue()
	}

	appsList, totalCount, err := s.provider.List(ctx, limit, req.GetOffset(), nameFilter, activeOnly)
	if err != nil {
		logger.Error("list operation failed", slog.String("error", err.Error()))
		return nil, status.Errorf(codes.Internal, "list operation failed")
	}

	logger.Debug("list completed", slog.Int("count", len(appsList)))
	return &apps.ListResponse{
		Data:       appsList,
		TotalCount: totalCount,
	}, nil
}
