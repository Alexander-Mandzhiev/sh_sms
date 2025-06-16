package subjects_handle

import (
	sl "backend/pkg/logger"
	library "backend/protos/gen/go/library"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"log/slog"
)

func (s *serverAPI) ListSubjects(ctx context.Context, _ *emptypb.Empty) (*library.ListSubjectsResponse, error) {
	const op = "grpc.PrivateSchool.Subjects.ListSubjects"
	logger := s.logger.With(slog.String("op", op))
	logger.Debug("List subjects called")

	subjects, err := s.service.ListSubjects(ctx)
	if err != nil {
		logger.Error("Failed to list subjects", sl.Err(err, true))
		return nil, s.convertError(err)
	}

	protoSubjects := make([]*library.Subject, 0, len(subjects))
	for _, subject := range subjects {
		protoSubjects = append(protoSubjects, subject.ToProto())
	}

	response := &library.ListSubjectsResponse{Subjects: protoSubjects}

	logger.Info("Subjects listed successfully", slog.Int("count", len(protoSubjects)))
	return response, nil
}
