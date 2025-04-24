package repository

import (
	"backend/service/constants"
	"backend/service/sso/models"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"log/slog"
)

func (r *Repository) Update(ctx context.Context, role *models.Role) (*models.Role, error) {
	const op = "repository.Roles.Update"
	logger := r.logger.With(slog.String("op", op), slog.String("role_id", role.ID.String()), slog.String("client_id", role.ClientID.String()))
	query := `UPDATE roles SET name = $1, description = $2, level = $3, is_active = $4, deleted_at = $5, updated_at = NOW() WHERE id = $6 AND client_id = $7 
			RETURNING id, client_id, name, description, level, is_custom, is_active, created_by, created_at, updated_at, deleted_at`
	args := []any{role.Name, role.Description, role.Level, role.IsActive, role.DeletedAt, role.ID, role.ClientID}

	var updated models.Role
	err := r.db.QueryRow(ctx, query, args...).Scan(&updated.ID, &updated.ClientID, &updated.Name, &updated.Description,
		&updated.Level, &updated.IsCustom, &updated.IsActive, &updated.CreatedBy, &updated.CreatedAt, &updated.UpdatedAt, &updated.DeletedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logger.Warn("role not found for update", slog.String("query", query), slog.Any("args", args))
			return nil, fmt.Errorf("%w: role", constants.ErrNotFound)
		}

		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				logger.Warn("unique constraint violation", slog.String("column", pgErr.ColumnName), slog.String("constraint", pgErr.ConstraintName))
				return nil, fmt.Errorf("%w: %s", constants.ErrConflict, pgErr.Detail)
			}
		}

		logger.Error("database operation failed", slog.String("query", query), slog.Any("arguments", args), slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", op, constants.ErrInternal)
	}

	logger.Debug("role successfully updated", slog.String("new_name", updated.Name), slog.Bool("is_active", updated.IsActive))
	return &updated, nil
}
