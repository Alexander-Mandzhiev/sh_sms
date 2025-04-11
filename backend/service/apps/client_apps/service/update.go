package service

import (
	sl "backend/pkg/logger"
	pb "backend/protos/gen/go/apps/clients_apps"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
)

func (s *Service) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.ClientApp, error) {
	const op = "service.Update"
	logger := s.logger.With(slog.String("op", op))
	if _, err := uuid.Parse(req.ClientId); err != nil {
		return nil, fmt.Errorf("%s: %w: invalid client_id", op, ErrInvalidArgument)
	}

	if req.AppId <= 0 {
		return nil, fmt.Errorf("%s: %w: invalid app_id", op, ErrInvalidArgument)
	}

	existing, err := s.provider.Get(ctx, req.ClientId, req.AppId)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil, fmt.Errorf("%s: %w", op, ErrNotFound)
		}
		logger.Error("failed to get existing record", sl.Err(err, true))
		return nil, fmt.Errorf("%s: %w", op, ErrInternal)
	}

	if req.IsActive != nil {
		existing.IsActive = req.GetIsActive()
	}

	existing.UpdatedAt = timestamppb.Now()
	if err = s.provider.Update(ctx, existing); err != nil {
		logger.Error("failed to update client app", sl.Err(err, true))
		return nil, fmt.Errorf("%s: %w", op, ErrInternal)
	}

	logger.Info("client app updated successfully")
	return existing, nil
}
