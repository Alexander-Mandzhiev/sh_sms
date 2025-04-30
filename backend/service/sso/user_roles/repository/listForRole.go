package repository

import (
	"backend/service/sso/models"
	"context"
	"fmt"
	"log/slog"
	"strings"
)

func (r *Repository) ListForRole(ctx context.Context, filter models.ListRequest) ([]*models.UserRole, int, error) {
	const op = "repository.UserRoles.ListForRole"
	logger := r.logger.With(slog.String("op", op), slog.Any("filter", filter))
	baseQuery := `SELECT user_id, role_id, client_id, app_id, assigned_by, expires_at, assigned_at FROM user_roles WHERE 1=1`
	countQuery := `SELECT COUNT(*) FROM user_roles WHERE 1=1`

	var (
		conditions []string
		args       []interface{}
		argNum     = 1
	)

	if filter.RoleID != nil {
		conditions = append(conditions, fmt.Sprintf("role_id = $%d", argNum))
		args = append(args, *filter.RoleID)
		argNum++
	}

	if filter.ClientID != nil {
		conditions = append(conditions, fmt.Sprintf("client_id = $%d", argNum))
		args = append(args, *filter.ClientID)
		argNum++
	}

	if filter.AppID != nil {
		conditions = append(conditions, fmt.Sprintf("app_id = $%d", argNum))
		args = append(args, *filter.AppID)
		argNum++
	}

	whereClause := ""
	if len(conditions) > 0 {
		whereClause = " AND " + strings.Join(conditions, " AND ")
	}

	total := 0
	if err := r.db.QueryRow(ctx, countQuery+whereClause, args...).Scan(&total); err != nil {
		logger.Error("count query failed", slog.Any("error", err))
		return nil, 0, fmt.Errorf("%s: %w", op, ErrInternal)
	}

	dataQuery := baseQuery + whereClause + " ORDER BY assigned_at DESC" + fmt.Sprintf(" LIMIT $%d OFFSET $%d", argNum, argNum+1)
	args = append(args, filter.Count, (filter.Page-1)*filter.Count)
	rows, err := r.db.Query(ctx, dataQuery, args...)
	if err != nil {
		logger.Error("data query failed", slog.Any("error", err))
		return nil, 0, fmt.Errorf("%s: %w", op, ErrInternal)
	}
	defer rows.Close()

	var result []*models.UserRole
	for rows.Next() {
		var ur models.UserRole
		if err = rows.Scan(&ur.UserID, &ur.RoleID, &ur.ClientID, &ur.AppID, &ur.AssignedBy, &ur.ExpiresAt, &ur.AssignedAt); err != nil {
			logger.Error("scan failed", slog.Any("error", err))
			return nil, 0, fmt.Errorf("%s: %w", op, ErrInternal)
		}
		result = append(result, &ur)
	}

	logger.Info("query executed successfully", slog.Int("returned_records", len(result)), slog.Int("total_records", total))
	return result, total, nil
}
