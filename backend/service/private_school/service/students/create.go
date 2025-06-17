package students_service

import (
	"backend/pkg/models/students"
	"backend/pkg/utils"
	"context"
	"log/slog"
)

func (s *Service) CreateStudent(ctx context.Context, st *students_models.CreateStudent) (*students_models.Student, error) {
	const op = "service.PrivateSchool.Students.CreateStudent"
	logger := s.logger.With(slog.String("op", op), slog.String("client_id", st.ClientID.String()), slog.String("contract", st.ContractNumber), slog.String("trace_id", utils.TraceIDFromContext(ctx)))
	logger.Debug("creating student")

	if ctx.Err() != nil {
		logger.Warn("request canceled before execution")
		return nil, ctx.Err()
	}

	student, err := s.provider.CreateStudent(ctx, st)
	if err != nil {
		return nil, s.handleRepoError(err, op, "full_name", st.FullName, "email", st.Email, "contract", st.ContractNumber)
	}

	logger.Info("student created", "student_id", student.ID)
	return student, nil
}
