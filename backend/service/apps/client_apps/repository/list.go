package repository

import (
	"backend/service/apps/constants"
	"backend/service/apps/models"
	"context"
	"fmt"
	"log/slog"
	"strings"
)

func (r *Repository) List(ctx context.Context, filter models.ListFilter) ([]*models.ClientApp, int, error) {
	const op = "repository.ClientApp.List"
	logger := r.logger.With(slog.String("op", op))

	baseQuery := `SELECT client_id, app_id, is_active, created_at, updated_at FROM client_apps WHERE is_active = true`
	countQuery := `SELECT COUNT(*) FROM client_apps WHERE is_active = true`

	var args []interface{}
	var conditions []string

	if filter.ClientID != nil && *filter.ClientID != "" {
		conditions = append(conditions, fmt.Sprintf("client_id = $%d", len(args)+1))
		args = append(args, *filter.ClientID)
	}

	if filter.AppID != nil && *filter.AppID > 0 {
		conditions = append(conditions, fmt.Sprintf("app_id = $%d", len(args)+1))
		args = append(args, *filter.AppID)
	}

	if filter.IsActive != nil {
		conditions = append(conditions, fmt.Sprintf("is_active = $%d", len(args)+1))
		args = append(args, *filter.IsActive)
	}

	if len(conditions) > 0 {
		whereClause := " AND " + strings.Join(conditions, " AND ")
		baseQuery += whereClause
		countQuery += whereClause
	}

	var total int
	if err := r.db.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		logger.Error("failed to get total count", slog.String("error", err.Error()))
		return nil, 0, fmt.Errorf("%s: %w", op, constants.ErrInternal)
	}

	baseQuery += fmt.Sprintf(" ORDER BY created_at DESC LIMIT $%d OFFSET $%d", len(args)+1, len(args)+2)
	args = append(args, filter.Count, (filter.Page-1)*filter.Count)

	rows, err := r.db.Query(ctx, baseQuery, args...)
	if err != nil {
		logger.Error("failed to list client apps", slog.String("error", err.Error()))
		return nil, 0, fmt.Errorf("%s: %w", op, constants.ErrInternal)
	}
	defer rows.Close()

	var apps []*models.ClientApp
	for rows.Next() {
		var app models.ClientApp
		if err = rows.Scan(&app.ClientID, &app.AppID, &app.IsActive, &app.CreatedAt, &app.UpdatedAt); err != nil {
			logger.Error("failed to scan row", slog.String("error", err.Error()))
			return nil, 0, fmt.Errorf("%s: %w", op, constants.ErrInternal)
		}
		apps = append(apps, &app)
	}

	if err = rows.Err(); err != nil {
		logger.Error("rows iteration error", slog.String("error", err.Error()))
		return nil, 0, fmt.Errorf("%s: %w", op, constants.ErrInternal)
	}

	logger.Debug("client apps listed successfully", slog.Int("count", len(apps)), slog.Int("total", total))
	return apps, total, nil
}
