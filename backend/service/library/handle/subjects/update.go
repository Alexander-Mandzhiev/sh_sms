package subjects_handle

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/library/subjects"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (s *serverAPI) Update(ctx context.Context, req *subjects.UpdateSubjectRequest) (*subjects.Subject, error) {
	s.logger.Debug("UpdateSubject called", slog.Any("request", req))

	resp, err := s.service.Update(ctx, req)
	if err != nil {
		s.logger.Error("Failed to update subject", sl.Err(err, true))
		return nil, status.Errorf(codes.InvalidArgument, "update failed: %v", err)
	}

	s.logger.Info("Subject updated", slog.Int("id", int(resp.Id)))
	return resp, nil
}
