package repository

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"backend/service/sso/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (r *Repository) GetByID(ctx context.Context, id uuid.UUID, appID int) (*models.Permission, error) {
	const op = "repository.Permission.GetByID"
	logger := r.logger.With(slog.String("op", op), slog.String("permission_id", id.String()), slog.Int("app_id", appID))
	query := `SELECT id, code, description, category, app_id, is_active, created_at, updated_at, deleted_at FROM permissions WHERE id = $1 AND app_id = $2`
	logger.Debug("executing query", slog.String("query", query), slog.Any("id", id), slog.Int("app_id", appID))
	row := r.db.QueryRow(ctx, query, id, appID)
	var perm models.Permission
	err := row.Scan(&perm.ID, &perm.Code, &perm.Description, &perm.Category, &perm.AppID, &perm.IsActive, &perm.CreatedAt, &perm.UpdatedAt, &perm.DeletedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logger.Warn("permission not found")
			return nil, fmt.Errorf("%w: %s", ErrNotFound, op)
		}
		logger.Error("database error", slog.Any("error", err), slog.String("id", id.String()))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	logger.Debug("permission found", slog.String("code", perm.Code), slog.Bool("is_active", perm.IsActive))
	return &perm, nil
}
