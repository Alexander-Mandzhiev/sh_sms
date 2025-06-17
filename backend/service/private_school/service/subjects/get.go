package subjects_service

import (
	"backend/pkg/models/subject"
	"context"
	"log/slog"
)

func (s *Service) GetSubject(ctx context.Context, id int32) (*subjects_models.Subject, error) {
	const op = "service.PrivateSchool.Subjects.GetSubject"
	logger := s.logger.With(slog.String("op", op))
	logger.Debug("retrieving subject", slog.Int("id", int(id)))

	if id <= 0 {
		logger.Warn("invalid subject ID")
		return nil, subjects_models.ErrInvalidSubjectID
	}

	subject, err := s.provider.GetSubjectByID(ctx, id)
	if err != nil {
		logger.Error("failed to get subject", slog.String("error", err.Error()), slog.Int("id", int(id)))
		return nil, err
	}

	logger.Debug("subject retrieved successfully", slog.Int("id", int(id)), slog.String("name", subject.Name))
	return subject, nil
}
