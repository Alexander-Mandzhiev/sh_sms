package service

import (
	sl "backend/pkg/logger"
	pb "backend/protos/gen/go/apps/clients_apps"
	"context"
	"errors"
	"fmt"
	"log/slog"
)

func (s *Service) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.ClientApp, error) {
	const op = "service.Update"
	logger := s.logger.With(slog.String("op", op))

	if err := validateClientID(req.ClientId); err != nil {
		logger.Warn("invalid client_id", slog.String("client_id", req.ClientId), sl.Err(err, false))
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	if err := validateAppID(req.AppId); err != nil {
		logger.Warn("invalid app_id", slog.Int("app_id", int(req.AppId)), sl.Err(err, false))
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	if req.IsActive == nil {
		logger.Warn("missing is_active field")
		return nil, fmt.Errorf("%s: %w", op, ErrInvalidArgument)
	}

	updatedApp, err := s.provider.Update(ctx, req.ClientId, int(req.AppId), *req.IsActive)
	if err != nil {
		logger.Error("update failed", slog.Any("error", err))
		if errors.Is(err, ErrNotFound) {
			return nil, fmt.Errorf("%s: %w", op, ErrNotFound)
		}
		return nil, fmt.Errorf("%s: %w", op, ErrInternal)
	}

	return updatedApp, nil
}
