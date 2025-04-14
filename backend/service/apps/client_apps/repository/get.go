package repository

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

func (r *Repository) Get(ctx context.Context, clientID string, appID int) (*pb.ClientApp, error) {
	const op = "repository.Get"
	logger := r.logger.With(slog.String("op", op))
	query := `SELECT client_id, app_id, is_active, created_at, updated_at FROM client_apps WHERE client_id = $1 AND app_id = $2`

	var app pb.ClientApp
	var createdAt, updatedAt time.Time

	err := r.db.QueryRow(ctx, query, clientID, appID).Scan(&app.ClientId, &app.AppId, &app.IsActive, &createdAt, &updatedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logger.Warn("client app not found", slog.String("client_id", clientID), slog.Int("app_id", appID))
			return nil, fmt.Errorf("%s: %w", op, ErrNotFound)
		}
		logger.Error("failed to get client app", sl.Err(err, false))
		return nil, fmt.Errorf("%s: %w", op, ErrInternal)
	}

	app.CreatedAt = timestamppb.New(createdAt)
	app.UpdatedAt = timestamppb.New(updatedAt)

	logger.Debug("client app retrieved", slog.String("client_id", clientID), slog.Int("app_id", appID))
	return &app, nil
}
