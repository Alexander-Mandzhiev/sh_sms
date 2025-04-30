package repository

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5/pgconn"
)

func (r *Repository) Delete(ctx context.Context, id int, permanent bool) error {
	const op = "repository.ClientType.Delete"
	logger := r.logger.With(slog.String("op", op), slog.Int("id", id), slog.Bool("permanent", permanent))
	logger.Debug("attempting to delete client type")

	tx, err := r.db.Begin(ctx)
	if err != nil {
		logger.Error("failed to start transaction", slog.Any("error", err))
		return fmt.Errorf("%w: %v", ErrInternal, err)
	}
	defer tx.Rollback(ctx)

	exists, err := r.Exist(ctx, id)
	if err != nil {
		logger.Error("existence check failed", slog.Any("error", err))
		return fmt.Errorf("%w: %v", ErrInternal, err)
	}
	if !exists {
		logger.Warn("client type not found")
		return fmt.Errorf("%w", ErrNotFound)
	}

	if permanent {
		var hasDeps bool
		hasDeps, err = r.HasDependentClients(ctx, id)
		if err != nil {
			logger.Error("dependency check failed", slog.Any("error", err))
			return fmt.Errorf("%w: %v", ErrInternal, err)
		}
		if hasDeps {
			logger.Warn("client type has dependencies")
			return fmt.Errorf("%w: existing clients", ErrConflict)
		}
	}

	var result pgconn.CommandTag
	if permanent {
		query := `DELETE FROM client_types WHERE id = $1`
		result, err = tx.Exec(ctx, query, id)
	} else {
		query := `UPDATE client_types SET is_active = false WHERE id = $1`
		result, err = tx.Exec(ctx, query, id)
	}

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			logger.Error("database error", slog.String("code", pgErr.Code), slog.String("detail", pgErr.Detail))
		}
		return fmt.Errorf("%w: %v", ErrInternal, err)
	}

	if result.RowsAffected() == 0 {
		logger.Warn("no rows affected")
		return fmt.Errorf("%w", ErrNotFound)
	}

	if err = tx.Commit(ctx); err != nil {
		logger.Error("transaction commit failed", slog.Any("error", err))
		return fmt.Errorf("%w: %v", ErrInternal, err)
	}

	logger.Info("delete operation completed")
	return nil
}
