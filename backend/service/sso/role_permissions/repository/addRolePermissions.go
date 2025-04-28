package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"log/slog"
)

func (r *Repository) AddRolePermissions(ctx context.Context, roleID uuid.UUID, clientID uuid.UUID, permissionIDs []uuid.UUID) error {
	const op = "repository.RolePermissions.AddRolePermissions"
	logger := r.logger.With(slog.String("op", op), slog.String("clientID", clientID.String()), slog.String("roleID", roleID.String()), slog.Int("permissions_count", len(permissionIDs)))
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
        INNER JOIN permissions p ON p.id = input_perms.perm_id
        INNER JOIN roles r ON 
            r.id = $1 AND 
            r.client_id = $3
        WHERE p.is_active = TRUE
        ON CONFLICT (role_id, permission_id) DO NOTHING`

	result, err := r.db.Exec(ctx, query, roleID, uuids, clientID)
	if err != nil {
		logger.Error("failed to insert permissions", slog.Any("error", err))
		return fmt.Errorf("%s: %w", op, err)
	}

	rowsAffected := result.RowsAffected()
	logger.Info("permissions added", slog.Int64("rows_affected", rowsAffected), slog.Int("attempted", len(permissionIDs)))
	return nil
}
