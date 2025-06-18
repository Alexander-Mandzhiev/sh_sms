package students_service

import (
	"backend/pkg/utils"
	"context"
	"github.com/google/uuid"
	"log/slog"
)

func (s *Service) HardDeleteStudent(ctx context.Context, id, clientID uuid.UUID) error {
	const op = "service.PrivateSchool.Students.HardDeleteStudent"
	logger := s.logger.With(slog.String("op", op), slog.String("student_id", id.String()), slog.String("client_id", clientID.String()), slog.String("trace_id", utils.TraceIDFromContext(ctx)))
	logger.Debug("hard delete student")

	if ctx.Err() != nil {
		logger.Warn("request canceled before execution")
		return ctx.Err()
	}

	if err := s.provider.HardDeleteStudent(ctx, id, clientID); err != nil {
		return s.handleRepoError(err, op)
	}

	logger.Debug("hard deleted student successfully")
	return nil
}
