package teachers_service

import (
	"backend/pkg/models/teacher"
	"context"
	"log/slog"
)

func (s *Service) CreateTeacher(ctx context.Context, params *teachers_models.CreateTeacher) (*teachers_models.Teacher, error) {
	const op = "teachers_service.CreateTeacher"
	logger := s.logger.With(slog.String("op", op), slog.String("teacher_id", params.ID.String()), slog.String("client_id", params.ClientID.String()))
	logger.Debug("creating teacher")

	teacher, err := s.provider.CreateTeacher(ctx, params)
	if err != nil {
		return nil, s.handleRepoError(logger, err, "CreateTeacher")
	}

	logger.Info("teacher created successfully")
	return teacher, nil
}
