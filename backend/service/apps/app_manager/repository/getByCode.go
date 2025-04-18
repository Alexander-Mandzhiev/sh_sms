package repository

import (
	"backend/service/apps/models"
	"backend/service/constants"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log/slog"
)

func (r *Repository) GetByCode(ctx context.Context, code string) (*models.App, error) {
	const op = "repository.AppRepository.GetByCode"
	logger := r.logger.With(slog.String("op", op), slog.String("code", code))
	query := `SELECT id, code, name, description, is_active, version, created_at, updated_at FROM apps WHERE code = $1`
	var app models.App
	err := r.db.QueryRow(ctx, query, code).Scan(&app.ID, &app.Code, &app.Name, &app.Description, &app.IsActive, &app.Version, &app.CreatedAt, &app.UpdatedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logger.Debug("Application not found", slog.String("code", code))
			return nil, constants.ErrNotFound
		}
		logger.Error("Database query failed", slog.String("code", code), slog.String("error", err.Error()), slog.String("query", query))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	logger.Debug("Application found", slog.String("code", app.Code), slog.Int("version", app.Version))
	return &app, nil
}
