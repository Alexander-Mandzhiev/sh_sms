package subjects_handle

import (
	sl "backend/pkg/logger"
	private_school_models "backend/pkg/models/private_school"
	library "backend/protos/gen/go/library"
	"context"

	"log/slog"
)

func (s *serverAPI) GetSubject(ctx context.Context, req *library.GetSubjectRequest) (*library.Subject, error) {
	const op = "grpc.PrivateSchool.Subjects.GetSubject"
	logger := s.logger.With(slog.String("op", op))
	logger.Debug("Get subject called", slog.Int("id", int(req.GetId())))

	if req.GetId() <= 0 {
		err := private_school_models.ErrInvalidSubjectID
		logger.Warn("Invalid subject ID", sl.Err(err, true), slog.Int("id", int(req.GetId())))
		return nil, s.convertError(err)
	}

	subject, err := s.service.GetSubject(ctx, req.GetId())
	if err != nil {
		logger.Error("Failed to get subject", sl.Err(err, true), slog.Int("id", int(req.GetId())))
		return nil, s.convertError(err)
	}

	protoSubject := subject.ToProto()
	logger.Debug("Subject retrieved successfully", slog.Int("id", int(protoSubject.GetId())), slog.String("name", protoSubject.GetName()))
	return protoSubject, nil
}
