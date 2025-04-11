package repository

import (
	pb "backend/protos/gen/go/apps/app_manager"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
	"time"
)

func (r *Repository) Get(ctx context.Context, req *pb.GetRequest) (*pb.App, error) {
	const op = "repository.Get"
	logger := r.logger.With(slog.String("op", op))

	query := "SELECT id, code, name, description, is_active, created_at, updated_at FROM apps WHERE "
	args := []interface{}{}

	if req.GetId() != 0 {
		query += "id = $1"
		args = append(args, req.GetId())
	} else if req.GetCode() != "" {
		query += "code = $1"
		args = append(args, req.GetCode())
	} else {
		return nil, fmt.Errorf("no identifier provided")
	}

	var app pb.App
	var createdAt, updatedAt time.Time
	err := r.db.QueryRow(ctx, query, args...).Scan(&app.Id, &app.Code, &app.Name, &app.Description, &app.IsActive, &createdAt, &updatedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logger.Warn("Application not found", slog.Any("request", req))
			return nil, ErrNotFound
		}
		logger.Error("Failed to get app", slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	app.CreatedAt = timestamppb.New(createdAt)
	app.UpdatedAt = timestamppb.New(updatedAt)

	return &app, nil
}
