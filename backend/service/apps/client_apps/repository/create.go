package repository

import (
	sl "backend/pkg/logger"
	pb "backend/protos/gen/go/apps/clients_apps"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
	"time"
)

func (r *Repository) Create(ctx context.Context, clientID string, appID int, isActive bool) (*pb.ClientApp, error) {
	const op = "repository.Create"
	logger := r.logger.With(slog.String("op", op))
	query := `INSERT INTO client_apps (client_id, app_id, is_active, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)`

	now := time.Now().UTC()
	_, err := r.db.Exec(ctx, query, clientID, appID, isActive, now, now)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			logger.Error("client app already exists", slog.String("client_id", clientID), slog.Int("app_id", appID))
			return nil, fmt.Errorf("%s: %w", op, ErrAlreadyExists)
		}
		logger.Error("failed to create client app", sl.Err(err, false))
		return nil, fmt.Errorf("%s: %w", op, ErrInternal)
	}

	timeNow := timestamppb.New(now)

	clientApp := &pb.ClientApp{
		ClientId:  clientID,
		AppId:     int32(appID),
		IsActive:  isActive,
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
	}

	logger.Info("client app created", slog.String("client_id", clientID), slog.Int("app_id", appID))
	return clientApp, nil
}
