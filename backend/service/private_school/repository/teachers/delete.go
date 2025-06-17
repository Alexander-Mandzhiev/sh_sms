package teachers_repository

import (
	teachers_models "backend/pkg/models/teacher"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *Repository) DeleteTeacher(ctx context.Context, id, clientID uuid.UUID) error {
	const op = "repository.PrivateSchool.Teachers.DeleteTeacher"
	logger := r.logger.With(slog.String("op", op))
	logger.Debug("Deleting teacher in repository")

	query := `UPDATE teachers SET deleted_at = $3 WHERE id = $1 AND client_id = $2 AND deleted_at IS NULL RETURNING id`
	var deletedID uuid.UUID
	err := r.db.QueryRow(ctx, query, id, clientID, time.Now().UTC()).Scan(&deletedID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logger.Warn("teacher not found or already deleted", "teacher_id", id, "client_id", clientID)
			return teachers_models.ErrTeacherNotFound
		}

		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23503" {
				logger.Warn("teacher has active references", "detail", pgErr.Detail, "constraint", pgErr.ConstraintName)
				return teachers_models.ErrDeleteTeacherConflict
			}
		}
		logger.Error("failed to delete teacher", "error", err)
		return fmt.Errorf("failed to delete teacher: %w", err)
	}

	logger.Info("teacher deleted successfully", "teacher_id", id, "client_id", clientID)
	return nil
}
