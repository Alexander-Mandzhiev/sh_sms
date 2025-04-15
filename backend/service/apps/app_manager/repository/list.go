package repository

import (
	sl "backend/pkg/logger"
	"backend/service/apps/models"
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) List(ctx context.Context, filter models.ListFilter) ([]models.App, int, error) {
	const op = "repository.AppRepository.List"
	logger := r.logger.With(slog.String("op", op), slog.Int("page", filter.Page), slog.Int("count", filter.Count))
	offset := (filter.Page - 1) * filter.Count
	baseQuery := `SELECT id, code, name, description, is_active, version, created_at, updated_at FROM apps`
	countQuery := `SELECT COUNT(*) FROM apps`

	where := ""
	args := []interface{}{}
	argCounter := 1

	if filter.FilterActive != nil {
		where = " WHERE is_active = $1"
		args = append(args, *filter.FilterActive)
		argCounter++
	}

	dataQuery := fmt.Sprintf(`%s%s ORDER BY id ASC LIMIT $%d OFFSET $%d`, baseQuery, where, argCounter, argCounter+1)

	dataArgs := args
	dataArgs = append(dataArgs, filter.Count, offset)

	rows, err := r.db.Query(ctx, dataQuery, dataArgs...)
	if err != nil {
		logger.Error("Data query failed", slog.String("query", dataQuery), slog.Any("args", dataArgs), slog.String("error", err.Error()))
		return nil, 0, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	apps := make([]models.App, 0)
	for rows.Next() {
		var app models.App
		err = rows.Scan(&app.ID, &app.Code, &app.Name, &app.Description, &app.IsActive, &app.Version, &app.CreatedAt, &app.UpdatedAt)
		if err != nil {
			logger.Error("Row scan failed", sl.Err(err, true))
			return nil, 0, fmt.Errorf("%s: %w", op, err)
		}
		apps = append(apps, app)
	}

	if err = rows.Err(); err != nil {
		logger.Error("Rows iteration failed", sl.Err(err, true))
		return nil, 0, fmt.Errorf("%s: %w", op, err)
	}

	total := 0
	err = r.db.QueryRow(ctx, countQuery+where, args...).Scan(&total)
	if err != nil {
		logger.Error("Count query failed", slog.String("query", countQuery), slog.Any("args", args), slog.String("error", err.Error()))
		return nil, 0, fmt.Errorf("%s: %w", op, err)
	}

	logger.Debug("List query executed", slog.Int("returned", len(apps)), slog.Int("total", total))
	return apps, total, nil
}
