package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"log/slog"
)

func (r *Repository) Delete(ctx context.Context, clientID, roleID uuid.UUID, appID int, permanent bool) error {
	const op = "repository.Roles.Delete"
	logger := r.logger.With(slog.String("op", op), slog.String("client_id", clientID.String()), slog.String("role_id", "REDACTED"))

	tx, err := r.db.Begin(ctx)
	if err != nil {
		logger.Error("failed to start transaction", slog.Any("error", err))
		return fmt.Errorf("%s: %w", op, ErrInternal)
	}
	defer tx.Rollback(ctx)

	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM roles WHERE id = $1 AND client_id = $2 AND app_id = $3)`
	err = tx.QueryRow(ctx, query, roleID, clientID, appID).Scan(&exists)
	if err != nil {
		logger.Error("existence check failed", slog.Any("error", err))
		return fmt.Errorf("%s: %w", op, ErrInternal)
	}
	if !exists {
		logger.Warn("role not found")
		return fmt.Errorf("%w", ErrNotFound)
	}

	queryDeps := `SELECT EXISTS(SELECT 1 FROM user_roles WHERE role_id = $1 AND expires_at IS NULL) OR EXISTS(SELECT 1 FROM role_permissions WHERE role_id = $1)`
	var hasDeps bool
	err = tx.QueryRow(ctx, queryDeps, roleID).Scan(&hasDeps)
	if err != nil {
		logger.Error("dependency check failed", slog.Any("error", err))
		return fmt.Errorf("%s: %w", op, ErrInternal)
	}
	if hasDeps {
		logger.Warn("role has active dependencies")
		return fmt.Errorf("%w: %v", ErrConflict, "active users assigned")
	}

	var result pgconn.CommandTag
	if permanent {
		queryDelete := `DELETE FROM roles WHERE id = $1 AND client_id = $2 AND app_id = $3`
		result, err = tx.Exec(ctx, queryDelete, roleID, clientID, appID)
	} else {
		queryDelete := `UPDATE roles SET deleted_at = NOW(), is_active = false WHERE id = $1 AND client_id = $2 AND app_id = $3`
		result, err = tx.Exec(ctx, queryDelete, roleID, clientID, appID)
	}

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23503" {
			logger.Warn("foreign key violation", slog.String("detail", pgErr.Detail))
			return fmt.Errorf("%w: %v", ErrConflict, "existing references")
		}
		logger.Error("delete operation failed", slog.Any("error", err))
		return fmt.Errorf("%s: %w", op, ErrInternal)
	}

	if result.RowsAffected() == 0 {
		logger.Warn("no rows affected")
		return fmt.Errorf("%w", ErrNotFound)
	}

	if err = tx.Commit(ctx); err != nil {
		logger.Error("transaction commit failed", slog.Any("error", err))
		return fmt.Errorf("%s: %w", op, ErrInternal)
	}

	logger.Info("role deleted successfully", slog.Bool("permanent", permanent))
	return nil
}
