package teachers_service

import (
	"backend/pkg/models/private_school"
	"context"
	"log/slog"
)

func (s *Service) UpdateTeacher(ctx context.Context, update *private_school_models.UpdateTeacher) (*private_school_models.Teacher, error) {
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
