package repository

import (
	"backend/service/sso/models"
	"context"
	"fmt"
	"log/slog"
	"strings"
)

func (r *Repository) List(ctx context.Context, req models.ListRequest) ([]models.Role, int, error) {
	const op = "repository.Roles.List"
	logger := r.logger.With(slog.String("op", op), slog.String("client_id", "REDACTED"))
	dataQuery, dataArgs := r.buildListQuery(req, true)

	roles, err := r.executeListQuery(ctx, dataQuery, dataArgs, logger)
	if err != nil {
		return nil, 0, err
	}

	countQuery, countArgs := r.buildListQuery(req, false)
	total, err := r.executeCountQuery(ctx, countQuery, countArgs, logger)
	if err != nil {
		return nil, 0, err
	}

	logger.Debug("roles listed", slog.Int("found", len(roles)), slog.Int("total", total))
	return roles, total, nil
}

func (r *Repository) buildListQuery(req models.ListRequest, includePagination bool) (string, []any) {
	baseQuery := "SELECT "
	if includePagination {
		baseQuery += `id, client_id, app_id, name, description, level, is_custom, is_active, created_by, created_at, updated_at, deleted_at FROM roles`
	} else {
		baseQuery += "COUNT(*) FROM roles"
	}

	query := baseQuery + " WHERE client_id = $1 AND app_id = $2"
	args := []any{req.ClientID, req.AppID}
	paramCounter := 3

	if req.NameFilter != nil && *req.NameFilter != "" {
		query += fmt.Sprintf(" AND name ILIKE $%d", paramCounter)
		args = append(args, "%"+strings.ToLower(*req.NameFilter)+"%")
		paramCounter++
	}

	if req.LevelFilter != nil {
		query += fmt.Sprintf(" AND level = $%d", paramCounter)
		args = append(args, *req.LevelFilter)
		paramCounter++
	}

	if req.ActiveOnly != nil {
		query += fmt.Sprintf(" AND is_active = $%d", paramCounter)
		args = append(args, *req.ActiveOnly)
		paramCounter++
	}

	if includePagination {
		query += fmt.Sprintf(" ORDER BY created_at DESC LIMIT $%d OFFSET $%d", paramCounter, paramCounter+1)
		args = append(args, req.Count, (req.Page-1)*req.Count)
	}

	return query, args
}

func (r *Repository) executeListQuery(ctx context.Context, query string, args []any, logger *slog.Logger) ([]models.Role, error) {
	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		logger.Error("query failed", slog.String("type", "data"), slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", "repository.Roles.List", ErrInternal)
	}
	defer rows.Close()

	var roles []models.Role
	for rows.Next() {
		var role models.Role
		err = rows.Scan(&role.ID, &role.ClientID, &role.AppID, &role.Name, &role.Description, &role.Level, &role.IsCustom,
			&role.IsActive, &role.CreatedBy, &role.CreatedAt, &role.UpdatedAt, &role.DeletedAt)
		if err != nil {
			logger.Error("row scan failed", slog.Any("error", err))
			return nil, fmt.Errorf("%s: %w", "repository.Roles.List", ErrInternal)
		}
		roles = append(roles, role)
	}

	if err = rows.Err(); err != nil {
		logger.Error("rows iteration error", slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", "repository.Roles.List", ErrInternal)
	}

	return roles, nil
}

func (r *Repository) executeCountQuery(ctx context.Context, query string, args []any, logger *slog.Logger) (int, error) {
	var total int
	err := r.db.QueryRow(ctx, query, args...).Scan(&total)
	if err != nil {
		logger.Error("query failed", slog.String("type", "count"), slog.Any("error", err))
		return 0, fmt.Errorf("%s: %w", "repository.Roles.List", ErrInternal)
	}
	return total, nil
}
