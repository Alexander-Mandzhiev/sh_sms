package subjects_service

import (
	library_models "backend/pkg/models/library"
	"context"
	"errors"
	"log/slog"
)

func (s *Service) CreateSubject(ctx context.Context, params *library_models.CreateSubjectParams) (*library_models.Subject, error) {
	const op = "service.Library.Subjects.CreateSubject"
	logger := s.logger.With(slog.String("op", op))
	logger.Debug("creating subject", slog.String("name", params.Name))

	subject := &library_models.Subject{
		Name: params.Name,
	}

	id, err := s.provider.CreateSubject(ctx, subject)
	if err != nil {
		logger.Error("failed to create subject in database", slog.String("error", err.Error()), slog.String("name", params.Name))
		if errors.Is(err, library_models.ErrDuplicateName) {
			logger.Warn("subject name already exists", slog.String("name", params.Name))
		}
		return nil, err
	}

	subject.ID = id
	logger.Info("subject created successfully", slog.Int("id", int(id)), slog.String("name", params.Name))
	return subject, nil
}
