package service

import (
	"backend/protos/gen/go/apps"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (s *Service) GetKeyRotationHistory(ctx context.Context, req *apps.GetKeyRotationHistoryRequest) (*apps.KeyRotationHistoryResponse, error) {
	const op = "service.GetKeyRotationHistory"
	logger := s.logger.With(slog.String("op", op), slog.String("app_id", req.GetAppId()))

	if req.GetAppId() == "" {
		logger.Error("empty app_id")
		return nil, status.Error(codes.InvalidArgument, "app_id is required")
	}

	limit := int32(100)
	if req.GetLimit() != nil {
		if l := req.GetLimit().GetValue(); l > 0 {
			if l > 1000 {
				logger.Warn("capping limit to 1000", slog.Int("requested", int(l)))
				l = 1000
			}
			limit = l
		} else {
			logger.Error("invalid limit", slog.Int("value", int(l)))
			return nil, status.Error(codes.InvalidArgument, "limit must be positive")
		}
	}

	records, err := s.provider.GetKeyRotationHistory(ctx, req.GetAppId(), limit)
	if err != nil {
		logger.Error("history retrieval failed", slog.String("error", err.Error()))
		return nil, status.Errorf(codes.Internal, "history retrieval failed")
	}

	logger.Debug("history retrieved", slog.Int("records", len(records)))
	return &apps.KeyRotationHistoryResponse{Data: records}, nil
}
