package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"log/slog"
	"time"
)

func (r *Repository) SoftDelete(ctx context.Context, clientID, roleID uuid.UUID) error {
	const op = "repository.Roles.SoftDelete"
	logger := r.logger.With(slog.String("op", op), slog.String("role_id", roleID.String()), slog.String("client_id", clientID.String()))

	query := `UPDATE roles SET deleted_at = NOW(), is_active = false WHERE id = $1 AND client_id = $2 AND deleted_at IS NULL RETURNING id`
	var id uuid.UUID
	err := r.db.QueryRow(ctx, query, roleID, clientID).Scan(&id)
	if err != nil {
		if err == pgx.ErrNoRows {
			logger.Debug("role not found or already deleted")
			return fmt.Errorf("%w: role", ErrNotFound)
		}
		logger.Error("database operation failed", slog.String("query", query), slog.Any("error", err))
		return fmt.Errorf("%s: %w", op, ErrInternal)
	}
	logger.Debug("role soft-deleted", slog.Time("deleted_at", time.Now()), slog.String("client_id", id.String()))
	return nil
}
