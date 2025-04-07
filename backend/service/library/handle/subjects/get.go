package subjects_handle

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/library/subjects"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (s *serverAPI) Get(ctx context.Context, req *subjects.GetSubjectRequest) (*subjects.Subject, error) {
	s.logger.Debug("GetSubject called", slog.Int("id", int(req.Id)))

	resp, err := s.service.Get(ctx, req)
	if err != nil {
		s.logger.Error("Failed to get subject", sl.Err(err, true))
		return nil, status.Errorf(codes.NotFound, "subject not found: %v", err)
	}

	s.logger.Info("Subject retrieved", slog.Int("id", int(resp.Id)))
	return resp, nil
}
