package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"log/slog"
)

func (r *Repository) HasDependencies(ctx context.Context, roleID uuid.UUID) (bool, error) {
	const op = "repository.Roles.HasDependencies"
	logger := r.logger.With(slog.String("op", op), slog.String("role_id", roleID.String()))

	query := `SELECT EXISTS(
            SELECT 1 FROM user_roles 
            WHERE role_id = $1 AND expires_at IS NULL
            UNION ALL
            SELECT 1 FROM role_permissions 
            WHERE role_id = $1 AND deleted_at IS NULL)`

	var exists bool
	err := r.db.QueryRow(ctx, query, roleID).Scan(&exists)
	if err != nil {
		logger.Error("dependency check failed", slog.Any("error", err), slog.String("query", query))
		return false, fmt.Errorf("%s: %w", op, err)
	}

	logger.Debug("dependencies status", slog.Bool("exists", exists), slog.String("role_id", roleID.String()))
	return exists, nil
}
