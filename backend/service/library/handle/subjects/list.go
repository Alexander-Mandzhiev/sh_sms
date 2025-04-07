package subjects_handle

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/library/subjects"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (s *serverAPI) List(ctx context.Context, req *subjects.ListSubjectsRequest) (*subjects.ListSubjectsResponse, error) {
	s.logger.Debug("ListSubjects called", slog.Any("filters", req))

	resp, err := s.service.List(ctx, req)
	if err != nil {
		s.logger.Error("Failed to list subjects", sl.Err(err, true))
		return nil, status.Errorf(codes.Internal, "list failed: %v", err)
	}

	s.logger.Info("Subjects listed", slog.Int("count", len(resp.Items)))
	return resp, nil
}
