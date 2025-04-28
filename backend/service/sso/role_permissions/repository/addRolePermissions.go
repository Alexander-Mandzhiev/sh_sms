package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"log/slog"
)

func (r *Repository) AddRolePermissions(ctx context.Context, roleID uuid.UUID, clientID uuid.UUID, appID int, permissionIDs []uuid.UUID) error {
	const op = "repository.RolePermissions.AddRolePermissions"
	logger := r.logger.With(slog.String("op", op), slog.String("clientID", clientID.String()), slog.String("roleID", roleID.String()), slog.Int("appID", appID), slog.Int("permissions_count", len(permissionIDs)))
	logger.Debug("attempting to add permissions to role")

	if len(permissionIDs) == 0 {
		logger.Warn("empty permissions list")
		return fmt.Errorf("empty permissions list")
	}

	uuids := make([][]byte, 0, len(permissionIDs))
	for _, id := range permissionIDs {
		uuids = append(uuids, id[:])
	}

	query := `INSERT INTO role_permissions (role_id, permission_id)
        SELECT $1, p.id 
        FROM unnest($2::uuid[]) AS input_perms(perm_id)
        INNER JOIN permissions p ON 
            p.id = input_perms.perm_id AND
            p.app_id = $4
        INNER JOIN roles r ON 
            r.id = $1 AND 
            r.client_id = $3 AND
            r.app_id = $4
        WHERE p.is_active = TRUE
        ON CONFLICT (role_id, permission_id) DO NOTHING`

	result, err := r.db.Exec(ctx, query, roleID, uuids, clientID, appID)
	if err != nil {
		logger.Error("failed to insert permissions", slog.Any("error", err), slog.String("query", query))
		return fmt.Errorf("%s: %w", op, err)
	}

	rowsAffected := result.RowsAffected()
	logger.Info("permissions added successfully", slog.Int64("rows_affected", rowsAffected), slog.Int("attempted_count", len(permissionIDs)))

	if rowsAffected != int64(len(permissionIDs)) {
		logger.Warn("partial insertion detected", slog.Int64("missing_count", int64(len(permissionIDs))-rowsAffected))
	}

	return nil
}
