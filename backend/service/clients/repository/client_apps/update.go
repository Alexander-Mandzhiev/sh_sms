package client_apps_repository

import (
	sl "backend/pkg/logger"
	pb "backend/protos/gen/go/apps/clients_apps"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
	"time"
)

func (r *Repository) Update(ctx context.Context, clientApp *pb.ClientApp) error {
	const op = "repository.Update"
	logger := r.logger.With(slog.String("op", op))
	query := `UPDATE client_apps SET is_active = $1, updated_at = $2 WHERE client_id = $3 AND app_id = $4 RETURNING updated_at`

	var updatedAt time.Time
	err := r.db.QueryRow(ctx, query, clientApp.IsActive, time.Now(), clientApp.ClientId, clientApp.AppId).Scan(&updatedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrNotFound
		}
		logger.Error("failed to update client app", sl.Err(err, true))
		return fmt.Errorf("%s: %w", op, ErrInternal)
	}

	clientApp.UpdatedAt = timestamppb.New(updatedAt)
	return nil
}
