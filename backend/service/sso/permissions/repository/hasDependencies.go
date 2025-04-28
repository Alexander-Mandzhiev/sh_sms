package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"log/slog"
)

func (r *Repository) HasDependencies(ctx context.Context, permissionID uuid.UUID) (bool, error) {
	const op = "repository.Permission.HasDependencies"
	logger := r.logger.With(slog.String("op", op), slog.String("permission_id", permissionID.String()))

	query := `SELECT EXISTS(SELECT 1 FROM role_permissions WHERE permission_id = $1 AND deleted_at IS NULL)`

	var exists bool
	err := r.db.QueryRow(ctx, query, permissionID).Scan(&exists)

	if err != nil {
		logger.Error("dependency check failed", slog.Any("error", err))
		return false, fmt.Errorf("%w: %v", ErrDatabase, err)
	}

	logger.Debug("dependencies status", slog.Bool("exists", exists), slog.String("permission_id", permissionID.String()))
	return exists, nil
}
