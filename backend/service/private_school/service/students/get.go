package students_service

import (
	"backend/pkg/models/students"
	"backend/pkg/utils"
	"context"
	"github.com/google/uuid"
	"log/slog"
)

func (s *Service) GetStudent(ctx context.Context, id, clientID uuid.UUID) (*students_models.Student, error) {
	const op = "service.PrivateSchool.Students.GetStudent"
	logger := s.logger.With(slog.String("op", op), slog.String("student_id", id.String()), slog.String("client_id", clientID.String()), slog.String("trace_id", utils.TraceIDFromContext(ctx)))
	logger.Debug("fetching student")

	if ctx.Err() != nil {
		logger.Warn("request canceled before execution")
		return nil, ctx.Err()
	}

	student, err := s.provider.GetStudent(ctx, id, clientID)
	if err != nil {
		return nil, s.handleRepoError(err, op)
	}

	logger.Debug("student retrieved successfully")
	return student, nil
}
