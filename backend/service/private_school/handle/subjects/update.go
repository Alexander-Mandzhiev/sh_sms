package subjects_handle

import (
	sl "backend/pkg/logger"
	private_school_models "backend/pkg/models/private_school"
	library "backend/protos/gen/go/library"
	"context"
	"log/slog"
)

func (s *serverAPI) UpdateSubject(ctx context.Context, req *library.UpdateSubjectRequest) (*library.Subject, error) {
	const op = "grpc.PrivateSchool.Subjects.UpdateSubject"
	logger := s.logger.With(slog.String("op", op))
	logger.Debug("Update subject called", slog.Int("id", int(req.GetId())), slog.String("name", req.GetName()))

	params, err := private_school_models.UpdateSubjectParamsFromProto(req)
	if err != nil {
		logger.Warn("Invalid update subject parameters", sl.Err(err, true), slog.Int("id", int(req.GetId())), slog.String("name", req.GetName()))
		return nil, s.convertError(err)
	}

	subject, err := s.service.UpdateSubject(ctx, params)
	if err != nil {
		logger.Error("Failed to update subject", sl.Err(err, true), slog.Int("id", int(params.ID)), slog.String("name", params.Name))
		return nil, s.convertError(err)
	}

	protoSubject := subject.ToProto()
	logger.Info("Subject updated successfully", slog.Int("id", int(protoSubject.GetId())), slog.String("name", protoSubject.GetName()))
	return protoSubject, nil
}
