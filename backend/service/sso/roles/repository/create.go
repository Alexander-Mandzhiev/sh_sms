package repository

import (
	"backend/service/constants"
	"backend/service/sso/models"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	"log/slog"
)

func (r *Repository) Create(ctx context.Context, role *models.Role) error {
	const op = "repository.Roles.CreateRole"
	logger := r.logger.With(slog.String("op", op), slog.String("role_id", role.ID.String()), slog.String("client_id", role.ClientID.String()))
	query := `INSERT INTO roles (id, client_id, name, description, level, is_custom, is_active, created_by, created_at, updated_at, deleted_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	_, err := r.db.Exec(ctx, query, role.ID, role.ClientID, role.Name, role.Description, role.Level, role.IsCustom, role.IsActive, role.CreatedBy, role.CreatedAt, role.UpdatedAt, role.DeletedAt)
	if err != nil {
		if isUniqueViolation(err) {
			logger.Warn("role name conflict", slog.String("name", role.Name))
			return fmt.Errorf("%w: %s", constants.ErrConflict, "role name already exists")
		}
		logger.Error("database error", slog.Any("error", err))
		return fmt.Errorf("%s: %w", op, constants.ErrInternal)
	}

	logger.Info("role created", slog.String("name", role.Name))
	return nil
}

func isUniqueViolation(err error) bool {
	var pgErr *pgconn.PgError
	return errors.As(err, &pgErr) && pgErr.Code == "23505"
}
