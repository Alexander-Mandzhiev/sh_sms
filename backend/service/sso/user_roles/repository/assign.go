package repository

import (
	"backend/service/sso/models"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	"log/slog"
)

func (r *Repository) Assign(ctx context.Context, role *models.UserRole) (*models.UserRole, error) {
	const op = "repository.UserRoles.Assign"
	logger := r.logger.With(slog.String("op", op), slog.String("user_id", role.UserID.String()), slog.String("role_id", role.RoleID.String()), slog.String("client_id", role.ClientID.String()))
	query := `INSERT INTO user_roles (user_id, role_id, client_id, app_id, assigned_by, expires_at, assigned_at) VALUES ($1, $2, $3, $4, $5, $6, $7)
        RETURNING user_id, role_id, client_id, app_id, assigned_by, expires_at, assigned_at`

	var assignedRole models.UserRole
	err := r.db.QueryRow(ctx, query, role.UserID, role.RoleID, role.ClientID, role.AppID, role.AssignedBy, role.ExpiresAt, role.AssignedAt).
		Scan(&assignedRole.UserID, &assignedRole.RoleID, &assignedRole.ClientID, &assignedRole.AppID,
			&assignedRole.AssignedBy, &assignedRole.ExpiresAt, &assignedRole.AssignedAt)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			logger.Warn("assignment conflict", slog.String("detail", pgErr.Detail))
			return nil, fmt.Errorf("%w: %s", ErrConflict, "role assignment exists")
		}

		if errors.As(err, &pgErr) && pgErr.Code == "23503" {
			logger.Error("foreign key violation", slog.String("constraint", pgErr.ConstraintName), slog.String("message", pgErr.Message))
			return nil, fmt.Errorf("%w: invalid user/role relation", ErrInvalidArgument)
		}

		logger.Error("database error", slog.Any("error", err), slog.String("query", query))
		return nil, fmt.Errorf("%s: %w", op, ErrInternal)
	}

	logger.Info("role assigned successfully", slog.Time("assigned_at", assignedRole.AssignedAt))
	return &assignedRole, nil
}
