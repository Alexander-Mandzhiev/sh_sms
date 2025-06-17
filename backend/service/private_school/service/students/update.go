package students_service

import (
	"backend/pkg/models/students"
	"backend/pkg/utils"
	"context"
	"log/slog"
)

func (s *Service) UpdateStudent(ctx context.Context, updateData *students_models.UpdateStudent) (*students_models.Student, error) {
	const op = "service.PrivateSchool.Students.UpdateStudent"
	logger := s.logger.With(slog.String("op", op), slog.String("student_id", updateData.ID.String()), slog.String("client_id", updateData.ClientID.String()), slog.String("trace_id", utils.TraceIDFromContext(ctx)))
	logger.Debug("updating student")

	if ctx.Err() != nil {
		logger.Warn("request canceled before execution")
		return nil, ctx.Err()
	}

	updatedStudent, err := s.provider.UpdateStudent(ctx, updateData)
	if err != nil {
		ct := []interface{}{}
		if updateData.FullName != nil {
			ct = append(ct, "full_name", *updateData.FullName)
		}
		if updateData.ContractNumber != nil {
			ct = append(ct, "contract", *updateData.ContractNumber)
		}

		return nil, s.handleRepoError(err, op, ct...)
	}

	logger.Info("student updated successfully", "student_id", updatedStudent.ID)
	return updatedStudent, nil
}
