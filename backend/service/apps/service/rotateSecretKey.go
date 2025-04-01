package service

import (
	"backend/protos/gen/go/apps"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
)

func (s *Service) RotateSecretKey(ctx context.Context, req *apps.RotateSecretKeyRequest) (*apps.SecretKeyResponse, error) {
	const op = "service.RotateSecretKey"
	logger := s.logger.With(slog.String("op", op), slog.String("app_id", req.GetAppId()))

	if req.GetAppId() == "" || req.GetRotatedBy() == "" {
		logger.Error("validation failed", slog.Bool("app_id_empty", req.GetAppId() == ""), slog.Bool("rotated_by_empty", req.GetRotatedBy() == ""))
		return nil, status.Error(codes.InvalidArgument, "app_id and rotated_by are required")
	}

	key, err := s.provider.RotateSecretKey(ctx, req.GetAppId(), req.GetRotatedBy(), req.GetInvalidatePrevious())
	if err != nil {
		logger.Error("rotation failed", slog.String("error", err.Error()))
		return nil, status.Errorf(codes.Internal, "rotation failed")
	}

	logger.Info("key rotated successfully")
	return &apps.SecretKeyResponse{
		AppId:       req.GetAppId(),
		SecretKey:   key,
		GeneratedAt: timestamppb.Now(),
	}, nil
}
