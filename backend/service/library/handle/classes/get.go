package classes_handle

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/library/classes"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (s *serverAPI) Get(ctx context.Context, req *classes.GetClassRequest) (*classes.Class, error) {
	s.logger.Debug("GetClass called", slog.Int("id", int(req.Id)))

	resp, err := s.service.Get(ctx, req)
	if err != nil {
		s.logger.Error("Failed to get class", sl.Err(err, true))
		return nil, status.Errorf(codes.NotFound, "class not found: %v", err)
	}

	s.logger.Info("Class retrieved", slog.Int("id", int(resp.Id)))
	return resp, nil
}
