package subjects_handle

import (
	sl "backend/pkg/logger"
	private_school_models "backend/pkg/models/private_school"
	library "backend/protos/gen/go/library"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"log/slog"
)

func (s *serverAPI) DeleteSubject(ctx context.Context, req *library.DeleteSubjectRequest) (*emptypb.Empty, error) {
	const op = "grpc.PrivateSchool.Subjects.DeleteSubject"
	logger := s.logger.With(slog.String("op", op))
	logger.Debug("Delete subject called", slog.Int("id", int(req.GetId())))

	if req.GetId() <= 0 {
		err := private_school_models.ErrInvalidSubjectID
		logger.Warn("Invalid subject ID for deletion", sl.Err(err, true), slog.Int("id", int(req.GetId())))
		return nil, s.convertError(err)
	}

	err := s.service.DeleteSubject(ctx, req.GetId())
	if err != nil {
		logger.Error("Failed to delete subject", sl.Err(err, true), slog.Int("id", int(req.GetId())))
		return nil, s.convertError(err)
	}

	logger.Info("Subject deleted successfully", slog.Int("id", int(req.GetId())))

	return &emptypb.Empty{}, nil
}
