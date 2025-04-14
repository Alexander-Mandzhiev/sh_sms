package repository

import (
	sl "backend/pkg/logger"
	pb "backend/protos/gen/go/apps/clients_apps"
	"backend/service/apps/client_apps/service"
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
	"strings"
	"time"
)

func (r *Repository) List(ctx context.Context, filter service.Filter, page, count int) ([]*pb.ClientApp, int, error) {
	const op = "repository.List"
	logger := r.logger.With(slog.String("op", op))
	baseQuery := `SELECT client_id, app_id, is_active, created_at, updated_at FROM client_apps WHERE 1=1`
	countQuery := `SELECT COUNT(*) FROM client_apps WHERE 1=1`

	var args []interface{}
	var conditions []string

	if filter.ClientID != "" {
		conditions = append(conditions, fmt.Sprintf("client_id = $%d", len(args)+1))
		args = append(args, filter.ClientID)
	}
	if filter.AppID > 0 {
		conditions = append(conditions, fmt.Sprintf("app_id = $%d", len(args)+1))
		args = append(args, filter.AppID)
	}
	if filter.IsActive != nil {
		conditions = append(conditions, fmt.Sprintf("is_active = $%d", len(args)+1))
		args = append(args, *filter.IsActive)
	}

	if len(conditions) > 0 {
		whereClause := " AND " + strings.Join(conditions, " AND ")
		baseQuery += whereClause
		countQuery += whereClause
	}

	var total int
	err := r.db.QueryRow(ctx, countQuery, args...).Scan(&total)
	if err != nil {
		logger.Error("failed to get total count", sl.Err(err, false))
		return nil, 0, fmt.Errorf("%s: %w", op, ErrInternal)
	}

	baseQuery += fmt.Sprintf(" ORDER BY created_at DESC LIMIT $%d OFFSET $%d", len(args)+1, len(args)+2)
	args = append(args, count, (page-1)*count)

	rows, err := r.db.Query(ctx, baseQuery, args...)
	if err != nil {
		logger.Error("failed to list client apps", sl.Err(err, false))
		return nil, 0, fmt.Errorf("%s: %w", op, ErrInternal)
	}
	defer rows.Close()

	var apps []*pb.ClientApp
	for rows.Next() {
		var app pb.ClientApp
		var createdAt, updatedAt time.Time

		err = rows.Scan(&app.ClientId, &app.AppId, &app.IsActive, &createdAt, &updatedAt)
		if err != nil {
			logger.Error("failed to scan row", sl.Err(err, false))
			return nil, 0, fmt.Errorf("%s: %w", op, ErrInternal)
		}

		app.CreatedAt = timestamppb.New(createdAt)
		app.UpdatedAt = timestamppb.New(updatedAt)
		apps = append(apps, &app)
	}

	if err = rows.Err(); err != nil {
		logger.Error("rows iteration error", sl.Err(err, false))
		return nil, 0, fmt.Errorf("%s: %w", op, ErrInternal)
	}

	logger.Debug("listed client apps", slog.Int("count", len(apps)), slog.Int("total", total))
	return apps, total, nil
}
