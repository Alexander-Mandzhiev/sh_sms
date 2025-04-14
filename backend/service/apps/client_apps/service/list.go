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

	if err := validatePagination(req.Page, req.Count); err != nil {
		logger.Warn("invalid pagination", slog.Int("page", int(req.Page)), slog.Int("count", int(req.Count)), sl.Err(err, false))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if req.ClientId != nil {
		if err := validateClientID(*req.ClientId); err != nil {
			logger.Warn("invalid client_id filter", slog.String("client_id", *req.ClientId), sl.Err(err, false))
			return nil, fmt.Errorf("%s: %w", op, err)
		}
	}

	if req.AppId != nil {
		if err := validateAppID(*req.AppId); err != nil {
			logger.Warn("invalid app_id filter", slog.Int("app_id", int(*req.AppId)), sl.Err(err, false))
			return nil, fmt.Errorf("%s: %w", op, err)
		}
	}

	filter := Filter{
		ClientID: *req.ClientId,
		AppID:    int(*req.AppId),
		IsActive: req.IsActive,
	}

	apps, total, err := s.provider.List(ctx, filter, int(req.Page), int(req.Count))
	if err != nil {
		logger.Error("list failed", slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", op, ErrInternal)
	}

	return &pb.ListResponse{
		ClientApps: apps,
		TotalCount: int32(total),
		Page:       req.Page,
		Count:      req.Count,
	}, nil
}
