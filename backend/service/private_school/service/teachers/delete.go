package teachers_service

import (
	"backend/pkg/models/teacher"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log/slog"
)

func (s *Service) DeleteTeacher(ctx context.Context, id, clientID uuid.UUID) error {
	const op = "teachers_service.DeleteTeacher"
	logger := s.logger.With(slog.String("op", op), slog.String("teacher_id", id.String()), slog.String("client_id", clientID.String()))
	logger.Debug("deleting teacher")

	if err := s.provider.DeleteTeacher(ctx, id, clientID); err != nil {
		if errors.Is(err, teachers_models.ErrTeacherNotFound) {
			logger.Warn("teacher not found or already deleted", "error", err, "teacher_id", id, "client_id", clientID)
			return err
		}
		if errors.Is(err, teachers_models.ErrDeleteTeacherConflict) {
			logger.Warn("teacher has active references, cannot delete", "error", err, "teacher_id", id, "client_id", clientID)
			return err
		}
		logger.Error("failed to delete teacher", "error", err, "teacher_id", id, "client_id", clientID)
		return fmt.Errorf("failed to delete teacher: %w", err)
	}

	logger.Info("teacher deleted successfully")
	return nil
}
