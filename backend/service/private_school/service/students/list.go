package students_service

import (
	"backend/pkg/models/students"
	"backend/pkg/utils"
	"context"
	"log/slog"
)

func (s *Service) ListStudents(ctx context.Context, params *students_models.ListStudentsRequest) ([]*students_models.Student, string, error) {
	const op = "service.PrivateSchool.Students.ListStudents"
	logger := s.logger.With(slog.String("op", op), slog.String("client_id", params.ClientID.String()), slog.String("trace_id", utils.TraceIDFromContext(ctx)))
	logger.Debug("listing students", "count", params.Count, "cursor", params.Cursor)

	if ctx.Err() != nil {
		logger.Warn("request canceled before execution")
		return nil, "", ctx.Err()
	}

	if params.Count <= 0 || params.Count > students_models.MaxListCount {
		params.Count = 10
	}

	students, nextCursor, err := s.provider.ListStudents(ctx, params)
	if err != nil {
		return nil, "", s.handleRepoError(err, op)
	}

	logger.Debug("students listed successfully", "count", len(students))
	return students, nextCursor, nil
}
