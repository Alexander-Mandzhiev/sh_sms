package teachers_service

import (
	"backend/pkg/models/teacher"
	"context"
	"log/slog"
)

func (s *Service) UpdateTeacher(ctx context.Context, update *teachers_models.UpdateTeacher) (*teachers_models.Teacher, error) {
	const op = "teachers_service.UpdateTeacher"
	logger := s.logger.With(slog.String("op", op), slog.String("teacher_id", update.ID.String()), slog.String("client_id", update.ClientID.String()))
	logger.Debug("updating teacher")

	teacher, err := s.provider.UpdateTeacher(ctx, update)
	if err != nil {
		return nil, s.handleRepoError(logger, err, "UpdateTeacher")
	}

	logger.Info("teacher updated successfully")
	return teacher, nil
}
