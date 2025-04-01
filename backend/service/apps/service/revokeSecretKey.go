package service

import (
	"backend/protos/gen/go/apps"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"log/slog"
)

func (s *Service) RevokeSecretKey(ctx context.Context, req *apps.RevokeSecretKeyRequest) (*apps.RevokeSecretKeyResponse, error) {
	const op = "service.RevokeSecretKey"
	logger := s.logger.With(slog.String("op", op), slog.String("app_id", req.GetAppId()))

	if req.GetAppId() == "" || req.GetRevokedBy() == "" {
		logger.Error("validation failed", slog.Bool("app_id_empty", req.GetAppId() == ""), slog.Bool("revoked_by_empty", req.GetRevokedBy() == ""))
		return nil, status.Error(codes.InvalidArgument, "app_id and revoked_by are required")
	}

	newKey, err := s.provider.RevokeSecretKey(ctx, req.GetAppId(), req.GetRevokedBy(), req.GetRegenerate())
	if err != nil {
		logger.Error("failed to revoke secret key", slog.String("error", err.Error()))
		return nil, status.Errorf(codes.Internal, "failed to revoke secret key")
	}

	resp := &apps.RevokeSecretKeyResponse{Success: true, RevokedAt: timestamppb.Now()}
	if req.GetRegenerate() {
		resp.NewKey = wrapperspb.String(newKey)
	}

	logger.Info("secret key revoked", slog.Bool("regenerated", req.GetRegenerate()))
	return resp, nil
}
