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

func (r *Repository) Get(ctx context.Context, clientID string, appID int32) (*pb.ClientApp, error) {
	const op = "repository.Get"
	logger := r.logger.With(slog.String("op", op))

	query := `SELECT client_id, app_id, is_active, created_at, updated_at
		FROM client_apps
		WHERE client_id = $1 AND app_id = $2`

	var ca pb.ClientApp
	var createdAt, updatedAt time.Time

	err := r.db.QueryRow(ctx, query, clientID, appID).Scan(&ca.ClientId, &ca.AppId, &ca.IsActive, &createdAt, &updatedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		logger.Error("failed to get client app", sl.Err(err, true))
		return nil, fmt.Errorf("%s: %w", op, ErrInternal)
	}

	ca.CreatedAt = timestamppb.New(createdAt)
	ca.UpdatedAt = timestamppb.New(updatedAt)

	return &ca, nil
}
