package repository

import (
	"backend/service/apps/constants"
	"backend/service/apps/models"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log/slog"
)

func (r *Repository) GetByID(ctx context.Context, id int) (*models.App, error) {
	const op = "repository.AppRepository.GetByID"
	logger := r.logger.With(slog.String("op", op), slog.Int("id", id))
	query := `SELECT id, code, name, description, is_active, version, created_at, updated_at FROM apps WHERE id = $1`
	var app models.App
	err := r.db.QueryRow(ctx, query, id).Scan(&app.ID, &app.Code, &app.Name, &app.Description, &app.IsActive, &app.Version, &app.CreatedAt, &app.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logger.Debug("Application not found")
			return nil, constants.ErrNotFound
		}
		logger.Error("Database query failed", slog.String("error", err.Error()), slog.String("query", query))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	logger.Debug("Application found", slog.String("code", app.Code), slog.Bool("active", app.IsActive), slog.Int("version", app.Version))
	return &app, nil
}
