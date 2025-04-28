package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"log/slog"
)

func (r *Repository) ListRolesForPermission(ctx context.Context, permissionID uuid.UUID, clientID uuid.UUID, appID int) ([]uuid.UUID, error) {
	const op = "repository.RolePermissions.ListRolesForPermission"
	logger := r.logger.With(slog.String("op", op), slog.String("permissionID", permissionID.String()), slog.String("clientID", clientID.String()), slog.Int("appID", appID))
	logger.Debug("fetching roles for permission")

	query := `SELECT r.id FROM role_permissions rp
        JOIN roles r ON rp.role_id = r.id AND r.client_id = $2 AND r.app_id = $3 AND r.is_active = TRUE
        JOIN permissions p ON rp.permission_id = p.id AND p.id = $1 AND p.app_id = $3 AND p.is_active = TRUE
        WHERE rp.permission_id = $1`

	rows, err := r.db.Query(ctx, query, permissionID, clientID, appID)
	if err != nil {
		logger.Error("failed to execute query", slog.Any("error", err), slog.String("query", query))
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var roles []uuid.UUID
	for rows.Next() {
		var roleID uuid.UUID
		if err = rows.Scan(&roleID); err != nil {
			logger.Error("failed to scan role ID", slog.Any("error", err))
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		roles = append(roles, roleID)
	}

	if err = rows.Err(); err != nil {
		logger.Error("rows iteration error", slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	logger.Debug("successfully fetched roles", slog.Int("count", len(roles)))
	return roles, nil
}
