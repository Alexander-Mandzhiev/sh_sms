package service

import (
	"backend/protos/gen/go/apps"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
)

func (s *Service) GenerateSecretKey(ctx context.Context, req *apps.GenerateSecretKeyRequest) (*apps.SecretKeyResponse, error) {
	const op = "service.GenerateSecretKey"
	logger := s.logger.With(slog.String("op", op), slog.String("app_id", req.GetAppId()))
	if req.GetAppId() == "" || req.GetGeneratedBy() == "" {
		logger.Error("validation failed", slog.Bool("app_id_missing", req.GetAppId() == ""), slog.Bool("generated_by_missing", req.GetGeneratedBy() == ""))
		return nil, status.Error(codes.InvalidArgument, "app_id and generated_by are required")
	}

	keyLength := int32(32)
	if req.GetKeyLength() != nil {
		if l := req.GetKeyLength().GetValue(); l > 0 {
			keyLength = l
		} else {
			logger.Error("invalid key length", slog.Int("value", int(l)))
			return nil, status.Error(codes.InvalidArgument, "key_length must be positive")
		}
	}

	key, err := s.provider.GenerateSecretKey(ctx, req.GetAppId(), req.GetGeneratedBy(), keyLength)
	if err != nil {
		logger.Error("generation failed", slog.String("error", err.Error()))
		return nil, status.Errorf(codes.Internal, "generation failed")
	}

	logger.Info("key generated successfully")
	return &apps.SecretKeyResponse{
		AppId:       req.GetAppId(),
		SecretKey:   key,
		GeneratedAt: timestamppb.Now(),
	}, nil
}
