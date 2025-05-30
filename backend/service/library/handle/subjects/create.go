package subjects_handle

import (
	sl "backend/pkg/logger"
	"backend/pkg/models/library"
	library "backend/protos/gen/go/library"
	"context"
	"log/slog"
)

func (s *serverAPI) CreateSubject(ctx context.Context, req *library.CreateSubjectRequest) (*library.Subject, error) {
	const op = "grpc.Library.Subjects.CreateSubject"
	logger := s.logger.With(slog.String("op", op))
	logger.Debug("Create subject called", slog.String("name", req.GetName()))

	params, err := library_models.CreateSubjectParamsFromProto(req)
	if err != nil {
		logger.Warn("Invalid create subject parameters", sl.Err(err, true), slog.String("name", req.GetName()))
		return nil, s.convertError(err)
	}

	subject, err := s.service.CreateSubject(ctx, params)
	if err != nil {
		logger.Error("Failed to create subject", sl.Err(err, true), slog.String("name", params.Name))
		return nil, s.convertError(err)
	}

	protoSubject := subject.ToProto()
	logger.Info("Subject created successfully", slog.Int("id", int(protoSubject.GetId())), slog.String("name", protoSubject.GetName()))
	return protoSubject, nil
}
