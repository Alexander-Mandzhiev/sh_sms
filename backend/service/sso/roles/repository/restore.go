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

func (r *Repository) Restore(ctx context.Context, clientID, roleID uuid.UUID, appID int) (*models.Role, error) {
	const op = "repository.Roles.Restore"
	logger := r.logger.With(slog.String("op", op), slog.String("role_id", roleID.String()), slog.String("client_id", clientID.String()))
	query := `UPDATE roles SET deleted_at = NULL, is_active = TRUE WHERE id = $1 AND client_id = $2 AND app_id = $3 AND deleted_at IS NOT NULL 
		RETURNING id, client_id, app_id, name, description, level, is_custom, is_active, created_by, created_at, updated_at, deleted_at`

	var restored models.Role
	err := r.db.QueryRow(ctx, query, roleID, clientID, appID).Scan(
		&restored.ID, &restored.ClientID, &restored.AppID, &restored.Name, &restored.Description, &restored.Level, &restored.IsCustom,
		&restored.IsActive, &restored.CreatedBy, &restored.CreatedAt, &restored.UpdatedAt, &restored.DeletedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logger.Debug("role not found or not deleted", slog.String("query", query))
			return nil, fmt.Errorf("%w: role", ErrNotFound)
		}

		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				logger.Warn("unique constraint violation on restore", slog.String("column", pgErr.ColumnName), slog.String("constraint", pgErr.ConstraintName))
				return nil, fmt.Errorf("%w: %s", ErrConflict, pgErr.Detail)
			}
		}
		logger.Error("database operation failed", slog.String("query", query), slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", op, ErrInternal)
	}

	logger.Debug("role restored successfully", slog.Bool("is_active", restored.IsActive), slog.Time("updated_at", restored.UpdatedAt))
	return &restored, nil
}
