package repository

import (
	"backend/service/apps/models"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log/slog"
	"time"
)

func (r *Repository) Get(ctx context.Context, clientID string, appID int) (*models.ClientApp, error) {
	const op = "repository.ClientApp.Get"
	logger := r.logger.With(slog.String("op", op), slog.String("client_id", clientID), slog.Int("app_id", appID))
	query := `SELECT client_id, app_id, is_active, created_at, updated_at FROM client_apps WHERE client_id = $1 AND app_id = $2 AND is_active = true`

	var app models.ClientApp
	start := time.Now()
	defer func() {
		logger.Debug("query timing", slog.Duration("duration", time.Since(start)))
	}()

	if err := ctx.Err(); err != nil {
		logger.Warn("context cancelled before query")
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	err := r.db.QueryRow(ctx, query, clientID, appID).Scan(&app.ClientID, &app.AppID, &app.IsActive, &app.CreatedAt, &app.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logger.Warn("active client app not found")
			return nil, fmt.Errorf("%s: %w", op, ErrNotFound)
		}

		logger.Error("database query failed", slog.String("error", err.Error()))
		return nil, fmt.Errorf("%s: %w", op, ErrInternal)
	}

	if app.ClientID != clientID || app.AppID != appID {
		logger.Error("database returned inconsistent data", slog.String("returned_client_id", app.ClientID), slog.Int("returned_app_id", app.AppID))
		return nil, fmt.Errorf("%s: %w", op, ErrInternal)
	}

	logger.Info("client app retrieved successfully", slog.Bool("is_active", app.IsActive), slog.Time("updated_at", app.UpdatedAt))
	return &app, nil
}
