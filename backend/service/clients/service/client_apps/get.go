package client_apps_service

import (
	sl "backend/pkg/logger"
	pb "backend/protos/gen/go/apps/clients_apps"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log/slog"
)

func (s *Service) Get(ctx context.Context, req *pb.GetRequest) (*pb.ClientApp, error) {
	const op = "service.Get"
	logger := s.logger.With(slog.String("op", op))

	if _, err := uuid.Parse(req.ClientId); err != nil {
		return nil, fmt.Errorf("%s: %w: invalid client_id", op, ErrInvalidArgument)
	}

	if req.AppId <= 0 {
		return nil, fmt.Errorf("%s: %w: invalid app_id", op, ErrInvalidArgument)
	}

	clientApp, err := s.provider.Get(ctx, req.ClientId, req.AppId)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil, fmt.Errorf("%s: %w", op, ErrNotFound)
		}
		logger.Error("failed to get client app", sl.Err(err, true))
		return nil, fmt.Errorf("%s: %w", op, ErrInternal)
	}

	return clientApp, nil
}
