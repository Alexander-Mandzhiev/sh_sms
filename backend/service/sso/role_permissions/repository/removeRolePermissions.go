package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"log/slog"
)

func (r *Repository) RemoveRolePermissions(ctx context.Context, roleID uuid.UUID, clientID uuid.UUID, appID int, permissionIDs []uuid.UUID) (int, error) {
	const op = "repository.RolePermissions.RemoveRolePermissions"
	logger := r.logger.With(slog.String("op", op), slog.String("roleID", roleID.String()),
		slog.String("clientID", clientID.String()), slog.Int("appID", appID), slog.Int("permissions_count", len(permissionIDs)))
	logger.Debug("attempting to remove permissions from role")

	if len(permissionIDs) == 0 {
		logger.Warn("empty permissions list")
		return 0, fmt.Errorf("empty permissions list")
	}

	uuids := make([][]byte, 0, len(permissionIDs))
	for _, id := range permissionIDs {
		uuids = append(uuids, id[:])
	}

	query := `DELETE FROM role_permissions rp USING permissions p
        INNER JOIN roles r ON r.id = $1 AND r.client_id = $3 AND r.app_id = $4
        WHERE rp.role_id = $1 AND rp.permission_id = p.id AND p.app_id = $4 AND rp.permission_id = ANY($2::uuid[])`
	result, err := r.db.Exec(ctx, query, roleID, uuids, clientID, appID)
	if err != nil {
		logger.Error("failed to delete permissions", slog.Any("error", err), slog.String("query", query))
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	rowsAffected := result.RowsAffected()
	logger.Info("permissions removed", slog.Int64("rows_affected", rowsAffected), slog.Int("attempted_count", len(permissionIDs)))
	return int(rowsAffected), nil
}
