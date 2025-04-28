package repository

import (
	"backend/service/sso/models"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"log/slog"
)

func (r *Repository) GetByID(ctx context.Context, clientID, roleID uuid.UUID, appID int) (*models.Role, error) {
	const op = "repository.Roles.GetByID"
	logger := r.logger.With(slog.String("op", op), slog.String("client_id", clientID.String()), slog.String("role_id", roleID.String()))
	query := `SELECT id, client_id, app_id, name, description, level, is_custom, is_active, created_by, created_at, updated_at, deleted_at FROM roles 
    			WHERE id = $1 AND client_id = $2 AND app_id = $3 AND deleted_at IS NULL`
	var role models.Role
	err := r.db.QueryRow(ctx, query, roleID, clientID, appID).
		Scan(&role.ID, &role.ClientID, &role.AppID, &role.Name, &role.Description, &role.Level, &role.IsCustom,
			&role.IsActive, &role.CreatedBy, &role.CreatedAt, &role.UpdatedAt, &role.DeletedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logger.Warn("role not found")
			return nil, fmt.Errorf("%w", ErrNotFound)
		}
		logger.Error("database error", slog.Any("params", []interface{}{clientID, appID, roleID}), slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", op, ErrInternal)
	}

	logger.Debug("role retrieved successfully")
	return &role, nil
}
