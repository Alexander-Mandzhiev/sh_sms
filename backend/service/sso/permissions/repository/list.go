package repository

import (
	"backend/service/sso/models"
	"context"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5/pgconn"
)

func (r *Repository) List(ctx context.Context, filter models.ListRequest) ([]models.Permission, int, error) {
	const op = "repository.Permission.List"
	logger := r.logger.With(slog.String("op", op), slog.Int("app_id", *filter.AppID))
	baseQuery := `SELECT id, code, description, category, app_id, is_active, created_at, updated_at,	deleted_at FROM permissions`
	tail := " WHERE app_id = $1"
	args := []any{*filter.AppID}
	paramCounter := 2

	if filter.CodeFilter != nil && *filter.CodeFilter != "" {
		tail += fmt.Sprintf(" AND code ILIKE $%d", paramCounter)
		args = append(args, "%"+*filter.CodeFilter+"%")
		paramCounter++
	}

	if filter.CategoryFilter != nil && *filter.CategoryFilter != "" {
		tail += fmt.Sprintf(" AND category = $%d", paramCounter)
		args = append(args, *filter.CategoryFilter)
		paramCounter++
	}

	if filter.ActiveOnly != nil && *filter.ActiveOnly {
		tail += " AND is_active = true"
	}

	fullQuery := baseQuery + tail
	countQuery := "SELECT COUNT(*) FROM (" + fullQuery + ") AS subq"

	logger.Debug("executing count query", slog.String("query", countQuery))
	var total int
	if err := r.db.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			logger.Error("database error", slog.String("code", pgErr.Code), slog.String("message", pgErr.Message))
		}
		return nil, 0, fmt.Errorf("%s: %w", op, err)
	}

	dataQuery := fullQuery + " ORDER BY created_at DESC"
	dataQuery += fmt.Sprintf(" LIMIT $%d OFFSET $%d", paramCounter, paramCounter+1)
	args = append(args, filter.Count, (filter.Page-1)*filter.Count)

	logger.Debug("executing data query", slog.String("query", dataQuery), slog.Any("args", args))
	rows, err := r.db.Query(ctx, dataQuery, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var permissions []models.Permission
	for rows.Next() {
		var p models.Permission
		if err = rows.Scan(&p.ID, &p.Code, &p.Description, &p.Category, &p.AppID, &p.IsActive, &p.CreatedAt, &p.UpdatedAt, &p.DeletedAt); err != nil {
			logger.Error("row scan failed", slog.Any("error", err))
			return nil, 0, fmt.Errorf("%s: %w", op, err)
		}
		permissions = append(permissions, p)
	}

	if err = rows.Err(); err != nil {
		logger.Error("rows iteration error", slog.Any("error", err))
		return nil, 0, fmt.Errorf("%s: %w", op, err)
	}

	logger.Info("list executed successfully", slog.Int("results_count", len(permissions)), slog.Int("total_records", total))
	return permissions, total, nil
}
