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
	"time"
)

func (r *Repository) Create(ctx context.Context, params models.CreateClientApp) (*models.ClientApp, error) {
	const op = "repository.ClientApp.Create"
	logger := r.logger.With(slog.String("op", op))

	query := `INSERT INTO client_apps (client_id, app_id, is_active, created_at, updated_at) 
        VALUES ($1, $2, $3, $4, $5) RETURNING client_id, app_id, is_active, created_at, updated_at`

	now := time.Now().UTC()
	var app models.ClientApp
	err := r.db.QueryRow(ctx, query, params.ClientID, params.AppID, params.IsActive, now, now).Scan(&app.ClientID, &app.AppID, &app.IsActive, &app.CreatedAt, &app.UpdatedAt)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			logger.Error("client app already exists", slog.String("client_id", params.ClientID), slog.Int("app_id", params.AppID))
			return nil, fmt.Errorf("%s: %w", op, constants.ErrAlreadyExists)
		}

		logger.Error("failed to create client app", slog.String("error", err.Error()))
		return nil, fmt.Errorf("%s: %w", op, constants.ErrInternal)
	}

	logger.Info("client app created successfully", slog.String("client_id", params.ClientID), slog.Int("app_id", params.AppID))
	return &app, nil
}
