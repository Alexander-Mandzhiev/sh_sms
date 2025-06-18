package students_service

import (
	"backend/pkg/utils"
	"context"
	"github.com/google/uuid"
	"log/slog"
)

func (s *Service) RestoreStudent(ctx context.Context, id, clientID uuid.UUID) error {
	const op = "service.PrivateSchool.Students.RestoreStudent"
	logger := s.logger.With(slog.String("op", op), slog.String("student_id", id.String()), slog.String("client_id", clientID.String()), slog.String("trace_id", utils.TraceIDFromContext(ctx)))
	logger.Debug("restoring student")

	if ctx.Err() != nil {
		logger.Warn("request canceled before execution")
		return ctx.Err()
	}

	if err := s.provider.RestoreStudent(ctx, id, clientID); err != nil {
		return s.handleRepoError(err, op)
	}

	logger.Info("student restored successfully", "student_id", id.String(), "client_id", clientID.String())
	return nil
}
