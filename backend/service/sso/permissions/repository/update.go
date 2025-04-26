package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgerrcode"
	"log/slog"
	"time"

	"backend/service/sso/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *Repository) Update(ctx context.Context, perm models.Permission) (*models.Permission, error) {
	const op = "repository.Permission.Update"
	logger := r.logger.With(slog.String("op", op), slog.String("permission_id", perm.ID.String()), slog.Int("app_id", perm.AppID))
	logger.Debug("starting database update operation")

	query := `UPDATE permissions 
        SET code = $1, description = $2, category = $3, updated_at = NOW() WHERE id = $4 AND app_id = $5
        RETURNING id, code, description, category, app_id, is_active, created_at, updated_at, deleted_at`

	logger.Debug("executing SQL query", slog.String("query", query))
	row := r.db.QueryRow(ctx, query, perm.Code, perm.Description, perm.Category, perm.ID, perm.AppID)

	var updated models.Permission
	start := time.Now()
	err := row.Scan(&updated.ID, &updated.Code, &updated.Description, &updated.Category, &updated.AppID,
		&updated.IsActive, &updated.CreatedAt, &updated.UpdatedAt, &updated.DeletedAt)
	duration := time.Since(start)
	if err != nil {
		logger.Error("update operation failed", slog.Any("error", err), slog.Duration("duration", duration))
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == pgerrcode.UniqueViolation {
				logger.Warn("duplicate code violation", slog.String("constraint", pgErr.ConstraintName), slog.String("detail", pgErr.Detail))
				return nil, fmt.Errorf("%s: %w (detail: %s)", op, ErrAlreadyExists, pgErr.Detail)
			}

			if pgErr.Code == pgerrcode.CheckViolation {
				logger.Warn("check constraint violation", slog.String("constraint", pgErr.ConstraintName), slog.String("message", pgErr.Message))
				return nil, fmt.Errorf("%s: %w (constraint: %s)", op, ErrInvalidArgument, pgErr.ConstraintName)
			}
			logger.Error("unhandled database error", slog.String("code", pgErr.Code), slog.String("message", pgErr.Message))
			return nil, fmt.Errorf("%s: %w (code: %s)", op, ErrDatabase, pgErr.Code)
		}

		if errors.Is(err, pgx.ErrNoRows) {
			logger.Warn("permission not found")
			return nil, fmt.Errorf("%s: %w", op, ErrNotFound)
		}
		logger.Error("operation failed", slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	logger.Info("permission successfully updated", slog.Duration("duration", duration), slog.Time("new_updated_at", updated.UpdatedAt))
	return &updated, nil
}
