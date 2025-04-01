package repository

import (
	"backend/protos/gen/go/apps"
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
	"time"
)

func (r *Repository) Create(ctx context.Context, app *apps.App) (*apps.App, error) {
	const op = "repository.Create"
	logger := r.logger.With(slog.String("op", op))

	if r.db == nil {
		logger.Error(ErrNilConnection.Error())
		return nil, ErrNilConnection
	}

	query := `INSERT INTO apps (name, description, secret_key, created_by, metadata)
		VALUES ($1, $2, $3, $4, $5)	RETURNING id, created_at, updated_at`

	var createdAt, updatedAt time.Time
	err := r.db.QueryRow(ctx, query,
		app.Name,
		app.Description,
		app.SecretKey, // предполагается, что ключ уже зашифрован
		app.CreatedBy,
		app.Metadata,
	).Scan(&app.Id, &createdAt, &updatedAt)

	if err != nil {
		logger.Error("create failed", slog.String("error", err.Error()))
		return nil, fmt.Errorf("create failed: %w", err)
	}

	app.CreatedAt = timestamppb.New(createdAt)
	app.UpdatedAt = timestamppb.New(updatedAt)

	logger.Info("app created", slog.String("app_id", app.Id))
	return app, nil
}
