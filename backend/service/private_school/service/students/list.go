package students_service

import (
	"backend/pkg/models/students"
	"backend/pkg/utils"
	"context"
	"log/slog"
)

func (s *Service) ListStudents(ctx context.Context, params *students_models.ListStudentsRequest) ([]*students_models.Student, *students_models.Cursor, error) {
	const op = "service.Students.ListStudents"
	logger := s.logger.With(slog.String("op", op), slog.String("client_id", params.ClientID.String()), slog.String("trace_id", utils.TraceIDFromContext(ctx)))
	logger.Debug("listing students", "count", params.Count, "filter", params.Filter)

	students, nextCursor, err := s.provider.ListStudents(ctx, params)
	if err != nil {
		return nil, nil, s.handleRepoError(err, op)
	}

	logger.Debug("students listed successfully", "count", len(students))
	return students, nextCursor, nil
}
