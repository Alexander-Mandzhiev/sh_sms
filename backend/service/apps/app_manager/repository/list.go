package repository

import (
	pb "backend/protos/gen/go/apps/app_manager"
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
	"strings"
	"time"
)

func (r *Repository) List(ctx context.Context, req *pb.ListRequest) (*pb.ListResponse, error) {
	const op = "repository.List"
	logger := r.logger.With(slog.String("op", op))
	baseQuery := "SELECT id, code, name, description, is_active, created_at, updated_at FROM apps"
	countQuery := "SELECT COUNT(*) FROM apps"
	where := []string{}
	args := []interface{}{}

	if req.FilterIsActive != nil {
		where = append(where, fmt.Sprintf("is_active = $%d", len(args)+1))
		args = append(args, req.GetFilterIsActive())
	}

	if len(where) > 0 {
		clause := " WHERE " + strings.Join(where, " AND ")
		baseQuery += clause
		countQuery += clause
	}

	var total int32
	err := r.db.QueryRow(ctx, countQuery, args...).Scan(&total)
	if err != nil {
		logger.Error("Failed to count apps", slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	args = append(args, req.GetCount(), (req.GetPage()-1)*req.GetCount())
	baseQuery += fmt.Sprintf(" ORDER BY id LIMIT $%d OFFSET $%d", len(args)-1, len(args))

	rows, err := r.db.Query(ctx, baseQuery, args...)
	if err != nil {
		logger.Error("Failed to list apps", slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var apps []*pb.App
	for rows.Next() {
		var app pb.App
		var createdAt, updatedAt time.Time
		err = rows.Scan(&app.Id, &app.Code, &app.Name, &app.Description, &app.IsActive, &createdAt, &updatedAt)
		if err != nil {
			logger.Error("Failed to scan app row", slog.Any("error", err))
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		app.CreatedAt = timestamppb.New(createdAt)
		app.UpdatedAt = timestamppb.New(updatedAt)
		apps = append(apps, &app)
	}

	return &pb.ListResponse{
		Apps:       apps,
		TotalCount: total,
		Page:       req.GetPage(),
		Count:      int32(len(apps)),
	}, nil
}
