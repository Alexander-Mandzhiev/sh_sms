package repository

import (
	"backend/service/sso/models"
	"context"
	"fmt"
	"log/slog"
	"time"
)

func (r *Repository) ListForUser(ctx context.Context, filter models.ListRequest) ([]*models.UserRole, int, error) {
	const op = "repository.user_roles.ListForUser"
	logger := r.logger.With(slog.String("op", op), slog.String("user_id", filter.UserID.String()))
	dataQuery := `SELECT user_id, role_id, client_id, app_id, assigned_by, expires_at, assigned_at FROM user_roles WHERE user_id = $1 AND client_id = $2 AND app_id = $3`
	countQuery := `SELECT COUNT(*) FROM user_roles WHERE user_id = $1 AND client_id = $2 AND app_id = $3`
	params := []interface{}{filter.UserID, filter.ClientID, filter.AppID}

	activeFilter := ""
	if filter.ActiveOnly != nil && *filter.ActiveOnly {
		activeFilter = " AND (expires_at IS NULL OR expires_at > $4)"
		params = append(params, time.Now())
	}

	dataQuery += activeFilter + " ORDER BY assigned_at DESC LIMIT $4 OFFSET $5"
	countQuery += activeFilter

	limit := filter.Count
	offset := (filter.Page - 1) * filter.Count
	params = append(params, limit, offset)

	tx, err := r.db.Begin(ctx)
	if err != nil {
		logger.Error("transaction start failed", slog.Any("error", err))
		return nil, 0, fmt.Errorf("%s: %w", op, ErrInternal)
	}
	defer tx.Rollback(ctx)

	var totalCount int
	err = tx.QueryRow(ctx, countQuery, params[:len(params)-2]...).Scan(&totalCount)
	if err != nil {
		logger.Error("count query failed", slog.Any("error", err))
		return nil, 0, fmt.Errorf("%s: %w", op, ErrInternal)
	}

	rows, err := tx.Query(ctx, dataQuery, params...)
	if err != nil {
		logger.Error("data query failed", slog.Any("error", err), slog.Any("params", params))
		return nil, 0, fmt.Errorf("%s: %w", op, ErrInternal)
	}
	defer rows.Close()

	var roles []*models.UserRole
	for rows.Next() {
		var role models.UserRole
		err = rows.Scan(&role.UserID, &role.RoleID, &role.ClientID, &role.AppID, &role.AssignedBy, &role.ExpiresAt, &role.AssignedAt)
		if err != nil {
			logger.Error("row scan failed", slog.Any("error", err))
			return nil, 0, fmt.Errorf("%s: %w", op, ErrInternal)
		}
		roles = append(roles, &role)
	}

	if err = rows.Err(); err != nil {
		logger.Error("rows iteration error", slog.Any("error", err))
		return nil, 0, fmt.Errorf("%s: %w", op, ErrInternal)
	}

	if err = tx.Commit(ctx); err != nil {
		logger.Error("transaction commit failed", slog.Any("error", err))
		return nil, 0, fmt.Errorf("%s: %w", op, ErrInternal)
	}

	logger.Debug("query executed", slog.Int("returned", len(roles)), slog.Int("total", totalCount))
	return roles, totalCount, nil
}
