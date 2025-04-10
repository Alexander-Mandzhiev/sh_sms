package client_apps_repository

import (
	sl "backend/pkg/logger"
	pb "backend/protos/gen/go/apps/clients_apps"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"log/slog"
)

func (r *Repository) Create(ctx context.Context, clientApp *pb.ClientApp) error {
	const op = "repository.Create"
	logger := r.logger.With(slog.String("op", op))
	var pgErr *pgconn.PgError

	query := `INSERT INTO client_apps (client_id, app_id, is_active, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)`

	_, err := r.db.Exec(ctx, query, clientApp.ClientId, clientApp.AppId, clientApp.IsActive, clientApp.CreatedAt.AsTime(), clientApp.UpdatedAt.AsTime())
	if err != nil {
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			return ErrAlreadyExists
		}
		logger.Error("failed to create client app", sl.Err(err, true))
		return fmt.Errorf("%s: %w", op, ErrInternal)
	}

	return nil
}
