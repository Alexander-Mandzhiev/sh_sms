package classes_handle

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/library/classes"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (s *serverAPI) Update(ctx context.Context, req *classes.UpdateClassRequest) (*classes.Class, error) {
	s.logger.Debug("UpdateClass called", slog.Any("request", req))

	resp, err := s.service.Update(ctx, req)
	if err != nil {
		s.logger.Error("Failed to update class", sl.Err(err, true))
		return nil, status.Errorf(codes.InvalidArgument, "update failed: %v", err)
	}

	s.logger.Info("Class updated", slog.Int("id", int(resp.Id)))
	return resp, nil
}
