package students_repository

import (
	"backend/pkg/models/students"
	"backend/pkg/utils"
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *Repository) HardDeleteStudent(ctx context.Context, id, clientID uuid.UUID) error {
	const op = "repository.PrivateSchool.StudentsRepository.HardDeleteStudent"
	logger := r.logger.With(slog.String("op", op), slog.String("student_id", id.String()), slog.String("client_id", clientID.String()), slog.String("trace_id", utils.TraceIDFromContext(ctx)))
	logger.Debug("hard deleting student")

	if ctx.Err() != nil {
		logger.Warn("context canceled before execution")
		return ctx.Err()
	}

	query := `DELETE FROM students WHERE id = $1 AND client_id = $2`
	result, err := r.db.Exec(ctx, query, id, clientID)
	if err != nil {
		return r.handleHardDeleteError(err, op, logger)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		logger.Warn("student not found")
		return students_models.ErrStudentNotFound
	}

	logger.Info("student hard deleted", "rows_affected", rowsAffected)
	return nil
}

func (r *Repository) handleHardDeleteError(err error, op string, logger *slog.Logger) error {
	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		return err
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		logger.Error("database error", "code", pgErr.Code, "message", pgErr.Message, "detail", pgErr.Detail)
		if pgErr.Code == "23503" {
			logger.Warn("foreign key constraint violation")
			return students_models.ErrInvalidClientID
		}
		return fmt.Errorf("%w: %s", students_models.ErrDeleteFailed, pgErr.Message)

	}

	logger.Error("unexpected error", "error", err)
	return fmt.Errorf("%w: %v", students_models.ErrDeleteFailed, err)
}
