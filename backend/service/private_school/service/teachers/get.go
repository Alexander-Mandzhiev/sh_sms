package teachers_service

import (
	"backend/pkg/models/teacher"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log/slog"
)

func (s *Service) GetTeacher(ctx context.Context, id, clientID uuid.UUID) (*teachers_models.Teacher, error) {
	const op = "teachers_service.GetTeacher"
	logger := s.logger.With(slog.String("op", op), slog.String("teacher_id", id.String()), slog.String("client_id", clientID.String()))
	logger.Debug("getting teacher")

	teacher, err := s.provider.GetTeacher(ctx, id, clientID)
	if err != nil {
		if errors.Is(err, teachers_models.ErrTeacherNotFound) {
			logger.Warn("teacher not found")
			return nil, err
		}
		logger.Error("failed to get teacher", "error", err)
		return nil, fmt.Errorf("failed to get teacher: %w", err)
	}

	logger.Debug("teacher retrieved successfully")
	return teacher, nil
}
