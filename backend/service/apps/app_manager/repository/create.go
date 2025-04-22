package repository

import (
	"backend/service/apps/models"
	"backend/service/constants"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"log/slog"
)

func (r *Repository) Create(ctx context.Context, req *models.CreateApp) (*models.App, error) {
	const op = "repository.AppRepository.Create"
	logger := r.logger.With(slog.String("op", op), slog.String("code", req.Code))
	query := `INSERT INTO apps (code, name, description, is_active, created_at, updated_at)
        VALUES ($1, $2, $3, $4, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
        RETURNING id, code, name, description, is_active, created_at, updated_at`

	var app models.App
	err := r.db.QueryRow(ctx, query, req.Code, req.Name, req.Description, req.IsActive).
		Scan(&app.ID, &app.Code, &app.Name, &app.Description, &app.IsActive, &app.CreatedAt, &app.UpdatedAt)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			logger.Warn("Duplicate application code")
			return nil, fmt.Errorf("%w: code '%s'", constants.ErrAlreadyExists, req.Code)
		}
		logger.Error("Database error", slog.String("error", err.Error()), slog.String("code", req.Code))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	logger.Info("Application created", slog.Int("id", app.ID))
	return &app, nil
}
