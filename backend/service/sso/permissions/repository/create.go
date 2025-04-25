package repository

import (
	"backend/service/sso/models"
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *Repository) Create(ctx context.Context, perm models.Permission) (*models.Permission, error) {
	const op = "repository.Permission.Create"
	logger := r.logger.With(slog.String("op", op), slog.String("code", perm.Code), slog.Int("app_id", perm.AppID))
	logger.Debug("starting permission creation")

	if perm.Code == "" || perm.AppID <= 0 {
		logger.Warn("validation failed: missing required fields")
		return nil, fmt.Errorf("%s: %w", op, ErrInvalidArgument)
	}

	query := `INSERT INTO permissions (code, description, category, app_id, is_active) VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at, updated_at, deleted_at`
	logger.Debug("executing SQL query", slog.String("query", query))
	row := r.db.QueryRow(ctx, query, perm.Code, perm.Description, perm.Category, perm.AppID, perm.IsActive)

	var created models.Permission
	err := row.Scan(&created.ID, &created.CreatedAt, &created.UpdatedAt, &created.DeletedAt)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == pgerrcode.UniqueViolation {
				logger.Warn("duplicate code violation", slog.String("detail", pgErr.Detail))
				return nil, fmt.Errorf("%s: %w", op, ErrAlreadyExists)
			}
			logger.Error("database error", slog.String("code", pgErr.Code), slog.String("message", pgErr.Message))
			return nil, fmt.Errorf("%s: %w", op, ErrDatabase)
		}
		logger.Error("operation failed", slog.Any("error", err), slog.String("code", perm.Code))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	created.Code = perm.Code
	created.Description = perm.Description
	created.Category = perm.Category
	created.AppID = perm.AppID
	created.IsActive = perm.IsActive

	logger.Info("permission created", slog.String("new_id", created.ID.String()), slog.Time("created_at", created.CreatedAt))
	return &created, nil
}
