package service

import (
	sl "backend/pkg/logger"
	pb "backend/protos/gen/go/apps/clients_apps"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log/slog"
)

func (s *Service) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	const op = "service.Delete"
	logger := s.logger.With(slog.String("op", op))

	if _, err := uuid.Parse(req.ClientId); err != nil {
		return nil, fmt.Errorf("%s: %w: invalid client_id", op, ErrInvalidArgument)
	}

	if req.AppId <= 0 {
		return nil, fmt.Errorf("%s: %w: invalid app_id", op, ErrInvalidArgument)
	}

	if err := s.provider.Delete(ctx, req.ClientId, req.AppId); err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil, fmt.Errorf("%s: %w", op, ErrNotFound)
		}
		logger.Error("failed to delete client app", sl.Err(err, true))
		return nil, fmt.Errorf("%s: %w", op, ErrInternal)
	}

	logger.Info("client app deleted successfully")
	return &pb.DeleteResponse{Success: true}, nil
}
