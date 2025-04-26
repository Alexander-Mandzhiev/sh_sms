package repository

import (
	"backend/service/sso/models"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"log/slog"
)

func (r *Repository) Restore(ctx context.Context, id uuid.UUID, appID int) (*models.Permission, error) {
	const op = "repository.Permission.Restore"
	logger := r.logger.With(
		slog.String("op", op),
		slog.String("permission_id", id.String()),
		slog.Int("app_id", appID),
	)

	query := `UPDATE permissions 
		SET 
			is_active = true, 
			deleted_at = NULL,
			updated_at = NOW()
		WHERE id = $1 AND app_id = $2 AND deleted_at IS NOT NULL
		RETURNING 
			id, code, description, category, app_id, 
			is_active, created_at, updated_at, deleted_at`

	var perm models.Permission
	err := r.db.QueryRow(ctx, query, id, appID).Scan(&perm.ID, &perm.Code, &perm.Description, &perm.Category, &perm.AppID, &perm.IsActive, &perm.CreatedAt, &perm.UpdatedAt, &perm.DeletedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logger.Debug("permission not found or not deleted")
			return nil, fmt.Errorf("%w: %s", ErrNotFound, op)
		}

		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" { // Unique violation
				logger.Warn("unique constraint violation",
					slog.String("constraint", pgErr.ConstraintName),
					slog.String("detail", pgErr.Detail),
				)
				return nil, fmt.Errorf("%w: %s", ErrConflict, pgErr.Detail)
			}
		}

		logger.Error("database operation failed", slog.String("query", query), slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	logger.Info("permission restored successfully", slog.String("code", perm.Code), slog.Bool("is_active", perm.IsActive))
	return &perm, nil
}
