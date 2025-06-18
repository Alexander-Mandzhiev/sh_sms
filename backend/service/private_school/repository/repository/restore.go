package students_repository

import (
	"backend/pkg/models/students"
	"backend/pkg/utils"
	"context"
	"errors"
	"log/slog"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *Repository) RestoreStudent(ctx context.Context, id, clientID uuid.UUID) error {
	const op = "repository.PrivateSchool.StudentsRepository.RestoreStudent"
	logger := r.logger.With(slog.String("op", op), slog.String("student_id", id.String()), slog.String("client_id", clientID.String()), slog.String("trace_id", utils.TraceIDFromContext(ctx)))
	logger.Debug("restoring student")

	if ctx.Err() != nil {
		logger.Warn("context canceled before execution")
		return ctx.Err()
	}

	query := `UPDATE students SET deleted_at = NULL WHERE id = $1 AND client_id = $2 AND deleted_at IS NOT NULL`
	result, err := r.db.Exec(ctx, query, id, clientID)
	if err != nil {
		return r.handleRestoreError(err, op, logger)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		logger.Warn("student not found or not deleted")
		return students_models.ErrStudentNotDeleted
	}

	logger.Info("student restored", "rows_affected", rowsAffected, "student_id", id.String())
	return nil
}

func (r *Repository) handleRestoreError(err error, op string, logger *slog.Logger) error {
	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		return err
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		logger.Error("database error", "code", pgErr.Code, "message", pgErr.Message, "detail", pgErr.Detail, "constraint", pgErr.ConstraintName)

		if pgErr.Code == "23503" && pgErr.ConstraintName == "fk_client" {
			logger.Warn("foreign key constraint violation")
			return students_models.ErrInvalidClientID
		}
		if pgErr.Code == "23505" && pgErr.ConstraintName == "uniq_student_contract_active" {
			logger.Warn("unique constraint violation after restore")
			return students_models.ErrDuplicateContract
		}

		return students_models.ErrRestoreFailed

	}

	logger.Error("unexpected error", "error", err)
	return students_models.ErrRestoreFailed
}
