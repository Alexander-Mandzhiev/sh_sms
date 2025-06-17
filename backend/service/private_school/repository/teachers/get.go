package teachers_repository

import (
	"backend/pkg/models/teacher"
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (r *Repository) GetTeacher(ctx context.Context, id, clientID uuid.UUID) (*teachers_models.Teacher, error) {
	const op = "repository.PrivateSchool.Teachers.GetTeacher"
	logger := r.logger.With(slog.String("op", op))
	logger.Debug("Getting teacher from repository", "teacher_id", id, "client_id", clientID)

	query := `SELECT id, client_id, full_name, phone, email, additional_info, deleted_at, created_at, updated_at FROM teachers WHERE id = $1 AND client_id = $2`
	var t teachers_models.Teacher
	err := r.db.QueryRow(ctx, query, id, clientID).Scan(&t.ID, &t.ClientID, &t.FullName, &t.Phone, &t.Email, &t.AdditionalInfo, &t.DeletedAt, &t.CreatedAt, &t.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logger.Warn("teacher not found", "teacher_id", id, "client_id", clientID)
			return nil, teachers_models.ErrTeacherNotFound
		}
		logger.Error("failed to get teacher", "error", err, "teacher_id", id, "client_id", clientID)
		return nil, fmt.Errorf("failed to get teacher: %w", err)
	}

	logger.Debug("teacher retrieved successfully", "teacher_id", id, "client_id", clientID)
	return &t, nil
}
