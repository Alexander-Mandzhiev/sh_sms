package subjects_handle

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/library/subjects"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (s *serverAPI) Delete(ctx context.Context, req *subjects.DeleteSubjectRequest) (*subjects.DeleteSubjectResponse, error) {
	s.logger.Debug("DeleteSubject called", slog.Int("id", int(req.Id)))

	err := s.service.Delete(ctx, req)
	if err != nil {
		s.logger.Error("Failed to delete subject", sl.Err(err, true))
		return nil, status.Errorf(codes.Internal, "delete failed: %v", err)
	}

	s.logger.Info("Subject deleted", slog.Int("id", int(req.Id)))
	return &subjects.DeleteSubjectResponse{Success: true}, nil
}
