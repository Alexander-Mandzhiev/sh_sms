package repository

import (
	"backend/service/apps/constants"
	"backend/service/apps/models"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"log/slog"
	"time"
)

func (r *Repository) Update(ctx context.Context, params models.UpdateClientApp) (*models.ClientApp, error) {
	const op = "repository.ClientApp.Update"
	logger := r.logger.With(slog.String("op", op), slog.String("client_id", params.ClientID), slog.Int("app_id", params.AppID))

	if err := ctx.Err(); err != nil {
		logger.Warn("context cancelled before update")
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	query := `UPDATE client_apps SET is_active = COALESCE($1, is_active), updated_at = NOW() WHERE client_id = $2 AND app_id = $3
              RETURNING client_id, app_id, is_active, created_at, updated_at`

	var updatedApp models.ClientApp
	start := time.Now()
	err := r.db.QueryRow(ctx, query, params.IsActive, params.ClientID, params.AppID).Scan(
		&updatedApp.ClientID, &updatedApp.AppID, &updatedApp.IsActive, &updatedApp.CreatedAt, &updatedApp.UpdatedAt,
	)

	defer func() {
		logger.Debug("update query timing",
			slog.Duration("duration", time.Since(start)))
	}()

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logger.Warn("client app not found for update")
			return nil, fmt.Errorf("%s: %w", op, constants.ErrNotFound)
		}

		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			logger.Error("database error during update", slog.String("code", pgErr.Code), slog.String("message", pgErr.Message))
		} else {
			logger.Error("failed to update client app",
				slog.String("error", err.Error()))
		}
		return nil, fmt.Errorf("%s: %w", op, constants.ErrInternal)
	}

	logger.Info("client app updated successfully", slog.Bool("new_is_active", updatedApp.IsActive), slog.Time("updated_at", updatedApp.UpdatedAt))
	return &updatedApp, nil
}
