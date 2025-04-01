package service

import (
	"backend/protos/gen/go/apps"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
)

func (s *Service) Delete(ctx context.Context, req *apps.DeleteRequest) (*apps.DeleteResponse, error) {
	const op = "service.Delete"
	logger := s.logger.With(slog.String("op", op), slog.String("app_id", req.GetAppId()))
	if req.GetAppId() == "" || req.GetDeletedBy() == "" {
		logger.Error("validation failed", slog.Bool("app_id_missing", req.GetAppId() == ""), slog.Bool("deleted_by_missing", req.GetDeletedBy() == ""))
		return nil, status.Error(codes.InvalidArgument, "app_id and deleted_by are required")
	}

	err := s.provider.Delete(ctx, req.GetAppId())
	if err != nil {
		logger.Error("deletion failed", slog.String("error", err.Error()))
		return nil, status.Errorf(codes.Internal, "deletion failed")
	}

	logger.Info("app deleted successfully")
	return &apps.DeleteResponse{
		Success:   true,
		DeletedAt: timestamppb.Now(),
	}, nil
}
