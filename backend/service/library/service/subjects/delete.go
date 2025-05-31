package subjects_service

import (
	"backend/pkg/models/library"
	"context"
	"log/slog"
)

func (s *Service) DeleteSubject(ctx context.Context, id int32) error {
	const op = "service.Library.Subjects.DeleteSubject"
	logger := s.logger.With(slog.String("op", op))
	logger.Debug("deleting subject", slog.Int("id", int(id)))

	if id <= 0 {
		logger.Warn("invalid subject ID")
		return library_models.ErrInvalidSubjectID
	}

	err := s.provider.DeleteSubject(ctx, id)
	if err != nil {
		logger.Error("failed to delete subject", slog.String("error", err.Error()), slog.Int("id", int(id)))
		return err
	}

	logger.Info("subject deleted successfully", slog.Int("id", int(id)))
	return nil
}
