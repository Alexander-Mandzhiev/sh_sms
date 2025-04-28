package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"log/slog"
)

func (r *Repository) HasRelation(ctx context.Context, roleID uuid.UUID, permissionID uuid.UUID, clientID uuid.UUID, appID int) (bool, error) {
	const op = "repository.RolePermissions.HasRelation"
	logger := r.logger.With(slog.String("op", op), slog.String("roleID", roleID.String()), slog.String("permissionID", permissionID.String()), slog.String("clientID", clientID.String()), slog.Int("appID", appID))
	logger.Debug("checking role-permission relation")
	query := `SELECT EXISTS(SELECT 1 FROM role_permissions rp JOIN roles r ON rp.role_id = r.id AND r.id = $1 AND r.client_id = $3 AND r.app_id = $4 AND r.is_active = TRUE
            JOIN permissions p ON rp.permission_id = p.id AND p.id = $2 AND p.app_id = $4 AND p.is_active = TRUE)`

	var exists bool
	err := r.db.QueryRow(ctx, query, roleID, permissionID, clientID, appID).Scan(&exists)
	if err != nil {
		logger.Error("failed to check relation", slog.Any("error", err), slog.String("query", query))
		return false, fmt.Errorf("%s: %w", op, err)
	}

	logger.Debug("relation check result", slog.Bool("exists", exists))
	return exists, nil
}
