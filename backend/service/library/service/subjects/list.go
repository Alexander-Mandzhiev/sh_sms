package subjects_service

import (
	"backend/pkg/models/library"
	"context"
	"log/slog"
)

func (s *Service) ListSubjects(ctx context.Context) ([]*library_models.Subject, error) {
	const op = "service.Library.Subjects.ListSubjects"
	logger := s.logger.With(slog.String("op", op))
	logger.Debug("listing all subjects")

	subjects, err := s.provider.ListSubjects(ctx)
	if err != nil {
		logger.Error("failed to list subjects", slog.String("error", err.Error()))
		return nil, err
	}

	logger.Debug("subjects listed successfully", slog.Int("count", len(subjects)))
	return subjects, nil
}
