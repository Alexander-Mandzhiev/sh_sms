package client_apps_repository

import (
	sl "backend/pkg/logger"
	pb "backend/protos/gen/go/apps/clients_apps"
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
	"time"
)

func (r *Repository) List(ctx context.Context, filter *pb.ListRequest_Filter, limit, offset int32) ([]*pb.ClientApp, int32, error) {
	const op = "repository.List"
	logger := r.logger.With(slog.String("op", op))

	query := `SELECT client_id, app_id, is_active, created_at, updated_at FROM client_apps WHERE 1=1`
	countQuery := `SELECT COUNT(*)	FROM client_apps WHERE 1=1`

	args := []interface{}{}
	argCounter := 1

	if filter != nil {
		if filter.ClientId != "" {
			query += fmt.Sprintf(" AND client_id = $%d", argCounter)
			countQuery += fmt.Sprintf(" AND client_id = $%d", argCounter)
			args = append(args, filter.ClientId)
			argCounter++
		}

		if filter.AppId != 0 {
			query += fmt.Sprintf(" AND app_id = $%d", argCounter)
			countQuery += fmt.Sprintf(" AND app_id = $%d", argCounter)
			args = append(args, filter.AppId)
			argCounter++
		}

		if filter.IsActive != nil {
			query += fmt.Sprintf(" AND is_active = $%d", argCounter)
			countQuery += fmt.Sprintf(" AND is_active = $%d", argCounter)
			args = append(args, filter.GetIsActive())
			argCounter++
		}
	}

	var total int32
	err := r.db.QueryRow(ctx, countQuery, args...).Scan(&total)
	if err != nil {
		logger.Error("failed to get total count", sl.Err(err, true))
		return nil, 0, fmt.Errorf("%s: %w", op, ErrInternal)
	}

	query += " ORDER BY created_at DESC"
	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argCounter, argCounter+1)
	args = append(args, limit, offset)

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		logger.Error("failed to list client apps", sl.Err(err, true))
		return nil, 0, fmt.Errorf("%s: %w", op, ErrInternal)
	}
	defer rows.Close()

	var clientApps []*pb.ClientApp
	for rows.Next() {
		var ca pb.ClientApp
		var createdAt, updatedAt time.Time
		err = rows.Scan(&ca.ClientId, &ca.AppId, &ca.IsActive, &createdAt, &updatedAt)

		if err != nil {
			logger.Error("failed to scan row", sl.Err(err, true))
			return nil, 0, fmt.Errorf("%s: %w", op, ErrInternal)
		}

		ca.CreatedAt = timestamppb.New(createdAt)
		ca.UpdatedAt = timestamppb.New(updatedAt)
		clientApps = append(clientApps, &ca)
	}

	return clientApps, total, nil
}
