package subjects_handle

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/library/subjects"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (s *serverAPI) Create(ctx context.Context, req *subjects.CreateSubjectRequest) (*subjects.Subject, error) {
	s.logger.Debug("CreateSubject called", slog.Any("request", req))

	resp, err := s.service.Create(ctx, req)
	if err != nil {
		s.logger.Error("Failed to create subject", sl.Err(err, true))
		return nil, status.Errorf(codes.Internal, "failed to create subject: %v", err)
	}

	s.logger.Info("Subject created", slog.Int("id", int(resp.Id)))
	return resp, nil
}
