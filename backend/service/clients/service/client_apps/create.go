package client_apps_service

import (
	sl "backend/pkg/logger"
	pb "backend/protos/gen/go/apps/clients_apps"
	"context"
	"errors"
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
)

func (s *Service) Create(ctx context.Context, req *pb.CreateRequest) (*pb.ClientApp, error) {
	const op = "service.Create"
	logger := s.logger.With(slog.String("op", op))

	if err := validateCreateRequest(req); err != nil {
		logger.Error("validation failed", sl.Err(err, true))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	exists, err := s.provider.Get(ctx, req.ClientId, req.AppId)
	if err != nil && !errors.Is(err, ErrNotFound) {
		logger.Error("failed to check existence", sl.Err(err, true))
		return nil, fmt.Errorf("%s: %w", op, ErrInternal)
	}
	if exists != nil {
		return nil, fmt.Errorf("%s: %w", op, ErrAlreadyExists)
	}

	now := timestamppb.Now()

	clientApp := &pb.ClientApp{
		ClientId:  req.ClientId,
		AppId:     req.AppId,
		IsActive:  req.GetIsActive(),
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err = s.provider.Create(ctx, clientApp); err != nil {
		logger.Error("failed to create client app", sl.Err(err, true))
		return nil, fmt.Errorf("%s: %w", op, ErrInternal)
	}

	logger.Info("client app created successfully")
	return clientApp, nil
}
