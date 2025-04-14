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

func (r *Repository) Update(ctx context.Context, clientID string, appID int, isActive bool) (*pb.ClientApp, error) {
	const op = "repository.Update"
	logger := r.logger.With(slog.String("op", op))
	query := `UPDATE client_apps SET is_active = $1, updated_at = $2 WHERE client_id = $3 AND app_id = $4 RETURNING created_at, updated_at`

	now := time.Now().UTC()
	var createdAt, updatedAt time.Time

	err := r.db.QueryRow(ctx, query, isActive, now, clientID, appID).Scan(&createdAt, &updatedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logger.Warn("client app not found for update", slog.String("client_id", clientID), slog.Int("app_id", appID))
			return nil, fmt.Errorf("%s: %w", op, ErrNotFound)
		}
		logger.Error("failed to update client app", sl.Err(err, false))
		return nil, fmt.Errorf("%s: %w", op, ErrInternal)
	}

	logger.Info("client app updated", slog.String("client_id", clientID), slog.Int("app_id", appID), slog.Bool("is_active", isActive))

	return &pb.ClientApp{
		ClientId:  clientID,
		AppId:     int32(appID),
		IsActive:  isActive,
		CreatedAt: timestamppb.New(createdAt),
		UpdatedAt: timestamppb.New(updatedAt),
	}, nil
}
