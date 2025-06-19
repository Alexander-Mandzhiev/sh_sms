package groups_repository

import (
	"backend/pkg/models/groups"
	"backend/pkg/utils"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"log/slog"
)

func (r *Repository) DeleteGroup(ctx context.Context, publicID, clientID uuid.UUID) error {
	const op = "repository.PrivateSchool.GroupsRepository.DeleteGroup"
	logger := r.logger.With(slog.String("op", op), slog.String("public_id", publicID.String()), slog.String("client_id", clientID.String()), slog.String("trace_id", utils.TraceIDFromContext(ctx)))
	logger.Debug("deleting group from repository")

	query := `DELETE FROM groups WHERE public_id = $1 AND client_id = $2`

	if ctx.Err() != nil {
		logger.Warn("context canceled before query execution")
		return ctx.Err()
	}

	result, err := r.db.Exec(ctx, query, publicID, clientID)
	if err != nil {
		return r.handleDeleteError(err, query, logger)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		logger.Warn("group not found for deletion")
		return groups_models.ErrGroupNotFound
	}

	logger.Info("group deleted", "rows_affected", rowsAffected)
	return nil
}

func (r *Repository) handleDeleteError(err error, query string, logger *slog.Logger) error {
	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		logger.Warn("context error", "error", err)
		return err
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		if pgErr.Code == "23503" {
			constraint := pgErr.ConstraintName
			logger.Warn("foreign key violation during deletion", "constraint", constraint, "detail", pgErr.Detail)
			if constraint == "fk_group_in_some_related_table" {
				return groups_models.ErrDependentRecordsExist
			}
			return groups_models.ErrForeignKeyViolation
		}

		logger.Warn("database error during deletion", "code", pgErr.Code, "message", pgErr.Message, "detail", pgErr.Detail)
		return fmt.Errorf("%w: %s", groups_models.ErrDeleteFailed, pgErr.Message)
	}

	logger.Error("database operation failed", "error", err, "query", query)
	return fmt.Errorf("%w: %v", groups_models.ErrDeleteFailed, err)
}
