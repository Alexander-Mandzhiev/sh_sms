package classes_handle

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/library/classes"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (s *serverAPI) Create(ctx context.Context, req *classes.CreateClassRequest) (*classes.Class, error) {
	s.logger.Debug("CreateClass called", slog.Any("request", req))

	resp, err := s.service.Create(ctx, req)
	if err != nil {
		s.logger.Error("Failed to create class", sl.Err(err, true))
		return nil, status.Errorf(codes.Internal, "failed to create class: %v", err)
	}

	s.logger.Info("Class created", slog.Int("id", int(resp.Id)))
	return resp, nil
}
