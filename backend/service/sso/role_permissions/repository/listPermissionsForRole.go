package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"log/slog"
)

func (r *Repository) ListPermissionsForRole(
	ctx context.Context, roleID uuid.UUID, clientID uuid.UUID, appID int) ([]uuid.UUID, error) {
	const op = "repository.RolePermissions.ListPermissionsForRole"
	logger := r.logger.With(slog.String("op", op), slog.String("roleID", roleID.String()), slog.String("clientID", clientID.String()), slog.Int("appID", appID))
	logger.Debug("fetching permissions for role")

	query := `SELECT p.id
        FROM role_permissions rp
        JOIN permissions p ON rp.permission_id = p.id AND p.is_active = TRUE
        JOIN roles r ON rp.role_id = r.id AND r.id = $1 AND r.client_id = $2 AND r.app_id = $3 AND r.is_active = TRUE 
        WHERE rp.role_id = $1`

	rows, err := r.db.Query(ctx, query, roleID, clientID, appID)
	if err != nil {
		logger.Error("failed to execute query", slog.Any("error", err), slog.String("query", query))
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var permissions []uuid.UUID
	for rows.Next() {
		var permID uuid.UUID
		if err = rows.Scan(&permID); err != nil {
			logger.Error("failed to scan permission ID", slog.Any("error", err))
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		permissions = append(permissions, permID)
	}

	if err = rows.Err(); err != nil {
		logger.Error("rows iteration error", slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	logger.Debug("successfully fetched permissions", slog.Int("count", len(permissions)))
	return permissions, nil
}
