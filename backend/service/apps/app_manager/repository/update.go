package repository

import (
	"backend/service/apps/app_manager/handle"
	"backend/service/apps/models"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log/slog"
)

func (r *Repository) Update(ctx context.Context, app *models.App) (*models.App, error) {
	const op = "repository.AppRepository.Update"
	logger := r.logger.With(slog.String("op", op), slog.Int("app_id", app.ID), slog.Int("version", app.Version))
	query := `UPDATE apps SET code = $1, name = $2, description = $3, is_active = $4, version = version + 1, updated_at = CURRENT_TIMESTAMP
        WHERE id = $5 AND version = $6 RETURNING code, name, description, is_active, version, updated_at`

	err := r.db.QueryRow(ctx, query, app.Code, app.Name, app.Description, app.IsActive, app.ID, app.Version).
		Scan(&app.Code, &app.Name, &app.Description, &app.IsActive, &app.Version, &app.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logger.Warn("Version conflict or app not found")
			return nil, handle.ErrVersionConflict
		}
		logger.Error("Update failed", slog.String("query", query), slog.String("error", err.Error()))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	logger.Debug("App updated successfully", slog.Int("new_version", app.Version))
	return app, nil
}
