package subjects_service

import (
	private_school_models "backend/pkg/models/private_school"
	"context"
	"errors"
	"log/slog"
)

func (s *Service) CreateSubject(ctx context.Context, params *private_school_models.CreateSubjectParams) (*private_school_models.Subject, error) {
	const op = "service.PrivateSchool.Subjects.CreateSubject"
	logger := s.logger.With(slog.String("op", op))
	logger.Debug("creating subject", slog.String("name", params.Name))

	subject := &private_school_models.Subject{
		Name: params.Name,
	}

	id, err := s.provider.CreateSubject(ctx, subject)
	if err != nil {
		logger.Error("failed to create subject in database", slog.String("error", err.Error()), slog.String("name", params.Name))
		if errors.Is(err, private_school_models.ErrDuplicateSubjectName) {
			logger.Warn("subject name already exists", slog.String("name", params.Name))
		}
		return nil, err
	}

	subject.ID = id
	logger.Info("subject created successfully", slog.Int("id", int(id)), slog.String("name", params.Name))
	return subject, nil
}
