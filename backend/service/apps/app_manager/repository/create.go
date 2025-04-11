package repository

import (
	pb "backend/protos/gen/go/apps/app_manager"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"log/slog"
	"time"
)

func (r *Repository) Create(ctx context.Context, app *pb.App) (*pb.App, error) {
	const op = "repository.Create"
	logger := r.logger.With(slog.String("op", op))

	query := `INSERT INTO apps (code, name, description, is_active, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	now := time.Now().UTC()
	err := r.db.QueryRow(ctx, query, app.GetCode(), app.GetName(), app.GetDescription(), app.GetIsActive(), now, now).Scan(&app.Id)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			logger.Error("Application already exists", slog.String("code", app.GetCode()))
			return nil, ErrAlreadyExists
		}
		logger.Error("Failed to create app", slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	logger.Info("Application created", slog.Int("id", int(app.GetId())))
	return app, nil
}
