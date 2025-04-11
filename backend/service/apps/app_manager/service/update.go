package service

import (
	pb "backend/protos/gen/go/apps/app_manager"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
)

func (s *Service) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.App, error) {
	const op = "service.Update"
	logger := s.logger.With(slog.String("op", op))

	if req.GetId() <= 0 {
		logger.Error("Invalid ID")
		return nil, ErrInvalidID
	}

	if req.Name != nil && len(req.GetName()) > 250 {
		logger.Error("Name exceeds maximum length")
		return nil, status.Error(codes.InvalidArgument, "name must be ≤ 250 characters")
	}

	existingApp, err := s.provider.Get(ctx, &pb.GetRequest{Id: ptrInt32(req.GetId())})
	if err != nil {
		logger.Error("Failed to get app", slog.Int("id", int(req.GetId())), slog.Any("error", err))
		return nil, err
	}

	updated := false

	if req.Name != nil && req.GetName() != existingApp.GetName() {
		existingApp.Name = req.GetName()
		updated = true
	}

	if req.Description != nil && req.GetDescription() != existingApp.GetDescription() {
		existingApp.Description = req.GetDescription()
		updated = true
	}

	if req.IsActive != nil && req.GetIsActive() != existingApp.GetIsActive() {
		existingApp.IsActive = req.GetIsActive()
		updated = true
	}

	if req.Code != nil && req.GetCode() != existingApp.GetCode() {
		existingApp.Code = req.GetCode()
		updated = true
	}

	if !updated {
		logger.Error("No changes detected")
		return existingApp, nil
	}

	existingApp.UpdatedAt = timestamppb.Now()
	return s.provider.Update(ctx, existingApp)
}

func ptrInt32(v int32) *int32 {
	return &v
}
