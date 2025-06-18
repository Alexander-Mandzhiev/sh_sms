package students_repository

import (
	"backend/pkg/models/students"
	"backend/pkg/utils"
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"log/slog"
)

func (r *Repository) SoftDeleteStudent(ctx context.Context, id, clientID uuid.UUID) error {
	const op = "repository.PrivateSchool.StudentsRepository.SoftDeleteStudent"
	logger := r.logger.With(slog.String("op", op), slog.String("student_id", id.String()), slog.String("client_id", clientID.String()), slog.String("trace_id", utils.TraceIDFromContext(ctx)))
	logger.Debug("soft deleting student")

	if ctx.Err() != nil {
		logger.Warn("context canceled before execution")
		return ctx.Err()
	}

	query := `UPDATE students SET deleted_at = CURRENT_TIMESTAMP WHERE id = $1 AND client_id = $2 AND deleted_at IS NULL`
	result, err := r.db.Exec(ctx, query, id, clientID)
	if err != nil {
		return r.handleSoftDeleteError(err, op, logger)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		exists, err := r.studentExists(ctx, id, clientID)
		if err != nil {
			logger.Error("failed to check student existence", "error", err)
			return students_models.ErrInternal
		}
		if !exists {
			return students_models.ErrStudentNotFound
		}
		return students_models.ErrStudentAlreadyDeleted
	}

	logger.Info("student soft deleted", "rows_affected", rowsAffected, "student_id", id.String())
	return nil
}
func (r *Repository) handleSoftDeleteError(err error, op string, logger *slog.Logger) error {
	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		return err
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		logger.Error("database error", "code", pgErr.Code, "message", pgErr.Message, "detail", pgErr.Detail, "constraint", pgErr.ConstraintName)
		if pgErr.Code == "23503" && pgErr.ConstraintName == "fk_client" {
			return students_models.ErrInvalidClientID
		}
		return students_models.ErrDeleteFailed
	}

	logger.Error("unexpected error", "error", err)
	return students_models.ErrDeleteFailed
}
