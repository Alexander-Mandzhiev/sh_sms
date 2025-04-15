package service

import (
	sl "backend/pkg/logger"
	pb "backend/protos/gen/go/apps/clients_apps"
	"context"
	"errors"
	"fmt"
	"log/slog"
)

func (s *Service) Delete(ctx context.Context, req *pb.IdentifierRequest) (*pb.DeleteResponse, error) {
	const op = "service.Delete"
	logger := s.logger.With(slog.String("op", op))

	if err := validateClientID(req.ClientId); err != nil {
		logger.Warn("invalid client_id", slog.String("client_id", req.ClientId), sl.Err(err, false))
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	if err := validateAppID(req.AppId); err != nil {
		logger.Warn("invalid app_id", slog.Int("app_id", int(req.AppId)), sl.Err(err, false))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	err := s.provider.Delete(ctx, req.ClientId, int(req.AppId))
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			logger.Warn("client app not found", slog.Any("request", req))
			return nil, fmt.Errorf("%s: %w", op, ErrNotFound)
		}
		logger.Error("delete failed", slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", op, ErrInternal)
	}

	return &pb.DeleteResponse{Success: true}, nil
}
