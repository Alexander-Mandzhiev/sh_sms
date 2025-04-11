package repository

import (
	sl "backend/pkg/logger"
	pb "backend/protos/gen/go/apps/clients_apps"
	"context"
	"fmt"
	"log/slog"
	"time"
)

func (r *Repository) Update(ctx context.Context, clientApp *pb.ClientApp) error {
	const op = "repository.Update"
	logger := r.logger.With(slog.String("op", op))
	query := `UPDATE client_apps SET is_active = $1, updated_at = $2 WHERE client_id = $3 AND app_id = $4`
	result, err := r.db.Exec(ctx, query, clientApp.IsActive, time.Now(), clientApp.ClientId, clientApp.AppId)

	if err != nil {
		logger.Error("failed to update client app", sl.Err(err, true))
		return fmt.Errorf("%s: %w", op, ErrInternal)
	}

	if result.RowsAffected() == 0 {
		return ErrNotFound
	}

	return nil
}
