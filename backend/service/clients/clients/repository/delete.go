package repository

import (
	"backend/service/clients/clients/handle"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"log/slog"
)

func (r *Repository) Delete(ctx context.Context, id uuid.UUID, permanent bool) error {
	const op = "repository.Clients.Delete"
	logger := r.logger.With(slog.String("op", op), slog.String("client_id", id.String()), slog.Bool("permanent", permanent))

	var query string
	if permanent {
		query = `DELETE FROM clients WHERE id = $1`
	} else {
		query = `UPDATE clients SET is_active = FALSE, updated_at = NOW() WHERE id = $1 AND is_active = TRUE`
	}

	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23503" {
			logger.Warn("deletion conflict", slog.String("detail", pgErr.Detail))
			return handle.ErrDeletionConflict
		}
		logger.Error("database error", slog.String("error", err.Error()))
		return fmt.Errorf("%s: %w", op, handle.ErrInternal)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		logger.Debug("client not found or already inactive")
		return handle.ErrNotFound
	}

	logger.Debug("delete operation successful", slog.Int64("rows_affected", rowsAffected))
	return nil
}
