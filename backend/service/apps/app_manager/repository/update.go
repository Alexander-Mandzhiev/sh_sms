package repository

import (
	pb "backend/protos/gen/go/apps/app_manager"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (r *Repository) Update(ctx context.Context, app *pb.App) (*pb.App, error) {
	const op = "repository.Update"
	logger := r.logger.With(slog.String("op", op))
	query := `UPDATE apps SET code = $1, name = $2, description = $3, is_active = $4, updated_at = $5 WHERE id = $6`

	_, err := r.db.Exec(ctx, query, app.GetCode(), app.GetName(), app.GetDescription(), app.GetIsActive(), app.GetUpdatedAt().AsTime(), app.GetId())
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return nil, status.Error(codes.AlreadyExists, "code already exists")
		}

		logger.Error("update failed", slog.Int("id", int(app.GetId())), slog.String("error", err.Error()))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	logger.Info("Application updated", slog.Int("id", int(app.GetId())))
	return app, nil
}
