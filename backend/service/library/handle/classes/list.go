package classes_handle

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/library/classes"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (s *serverAPI) List(ctx context.Context, req *classes.ListClassesRequest) (*classes.ListClassesResponse, error) {
	s.logger.Debug("ListClasses called", slog.Any("filters", req))

	resp, err := s.service.List(ctx, req)
	if err != nil {
		s.logger.Error("Failed to list classes", sl.Err(err, true))
		return nil, status.Errorf(codes.Internal, "list failed: %v", err)
	}

	s.logger.Info("Classes listed", slog.Int("count", len(resp.Items)))
	return resp, nil
}
