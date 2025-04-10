package app_manager_handle

import (
	pb "backend/protos/gen/go/apps/app_manager"
	apps "backend/service/clients/service/apps_manager"
	"context"
	"log/slog"
)

func (s *serverAPI) List(ctx context.Context, req *pb.ListRequest) (*pb.ListResponse, error) {
	const op = "handler.List"
	logger := s.logger.With(slog.String("op", op))
	logger.Debug("Request received", slog.Int("page", int(req.Page)), slog.Int("count", int(req.Count)))

	if req.Page <= 0 || req.Count <= 0 || req.Count > 100 {
		logger.Error("Invalid pagination", slog.Int("page", int(req.Page)), slog.Int("count", int(req.Count)), slog.Int("max", 100))
		return nil, convertError(apps.ErrInvalidPagination)
	}

	res, err := s.service.List(ctx, req)
	if err != nil {
		logger.Error("List failed", slog.String("error", err.Error()))
		return nil, convertError(err)
	}

	logger.Debug("Results fetched", slog.Int("total", int(res.TotalCount)), slog.Int("returned", len(res.Apps)))
	return res, nil
}
