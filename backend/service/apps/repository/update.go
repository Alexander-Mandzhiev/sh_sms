package repository

import (
	"backend/protos/gen/go/apps"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
	"time"
)

func (r *Repository) Update(ctx context.Context, app *apps.App) (*apps.App, error) {
	const op = "repository.Update"
	logger := r.logger.With(slog.String("op", op), slog.String("app_id", app.Id))

	if r.db == nil {
		logger.Error(ErrNilConnection.Error())
		return nil, ErrNilConnection
	}

	query := `
		UPDATE apps 
		SET name = $2,
			description = $3,
			is_active = $4,
			metadata = $5,
			updated_at = NOW()
		WHERE id = $1
		RETURNING updated_at
	`

	var updatedAt time.Time
	err := r.db.QueryRow(ctx, query,
		app.Id,
		app.Name,
		app.Description,
		app.IsActive,
		app.Metadata,
	).Scan(&updatedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logger.Error("app not found")
			return nil, ErrAppNotFound
		}
		logger.Error("update failed", slog.String("error", err.Error()))
		return nil, fmt.Errorf("update failed: %w", err)
	}

	app.UpdatedAt = timestamppb.New(updatedAt)
	logger.Info("app updated")
	return app, nil
}
