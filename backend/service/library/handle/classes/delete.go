package classes_handle

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/library/classes"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (s *serverAPI) Delete(ctx context.Context, req *classes.DeleteClassRequest) (*classes.DeleteClassResponse, error) {
	s.logger.Debug("DeleteClass called", slog.Int("id", int(req.Id)))

	err := s.service.Delete(ctx, req)
	if err != nil {
		s.logger.Error("Failed to delete class", sl.Err(err, true))
		return nil, status.Errorf(codes.Internal, "delete failed: %v", err)
	}

	s.logger.Info("Class deleted", slog.Int("id", int(req.Id)))
	return &classes.DeleteClassResponse{Success: true}, nil
}
