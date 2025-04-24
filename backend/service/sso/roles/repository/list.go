package repository

import (
	"backend/service/constants"
	"backend/service/sso/models"
	"context"
	"fmt"
	"log/slog"
	"strings"
)

func (r *Repository) List(ctx context.Context, req models.ListRequest) ([]models.Role, int, error) {
	const op = "repository.Roles.List"
	logger := r.logger.With(slog.String("op", op), slog.String("client_id", req.ClientID.String()))
	dataQuery, dataArgs := r.buildDataQuery(req)

	roles, err := r.executeDataQuery(ctx, dataQuery, dataArgs, logger)
	if err != nil {
		return nil, 0, err
	}
	countQuery, countArgs := r.buildCountQuery(req)
	total, err := r.executeCountQuery(ctx, countQuery, countArgs, logger)
	if err != nil {
		return nil, 0, err
	}

	logger.Debug("roles listed successfully", slog.Int("found", len(roles)), slog.Int("total", total))
	return roles, total, nil
}

func (r *Repository) buildDataQuery(req models.ListRequest) (string, []any) {
	query := `SELECT id, client_id, name, description, level, is_custom, is_active, created_by, created_at, updated_at, deleted_at 
		FROM roles WHERE client_id = $1`
	paramCounter := 2
	args := []any{req.ClientID}

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
		args = append(args, req.ActiveOnly)
		paramCounter++
	}

	query += fmt.Sprintf(" ORDER BY created_at DESC LIMIT $%d OFFSET $%d", paramCounter, paramCounter+1)
	args = append(args, req.Count, (req.Page-1)*req.Count)

	return query, args
}

func (r *Repository) buildCountQuery(req models.ListRequest) (string, []any) {
	query := "SELECT COUNT(*) FROM roles WHERE client_id = $1"
	args := []any{req.ClientID}
	paramCounter := 2

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

	return query, args
}

func (r *Repository) executeDataQuery(ctx context.Context, query string, args []any, logger *slog.Logger) ([]models.Role, error) {
	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		logger.Error("data query failed", slog.String("query", query), slog.Any("args", args), slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", "repository.Roles.List", constants.ErrInternal)
	}
	defer rows.Close()

	var roles []models.Role
	for rows.Next() {
		var role models.Role
		if err = rows.Scan(&role.ID, &role.ClientID, &role.Name, &role.Description, &role.Level, &role.IsCustom, &role.IsActive, &role.CreatedBy, &role.CreatedAt, &role.UpdatedAt, &role.DeletedAt); err != nil {
			logger.Error("scan failed", slog.Any("error", err))
			return nil, fmt.Errorf("%s: %w", "repository.Roles.List", constants.ErrInternal)
		}
		roles = append(roles, role)
	}

	if err = rows.Err(); err != nil {
		logger.Error("rows iteration error", slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", "repository.Roles.List", constants.ErrInternal)
	}

	return roles, nil
}

func (r *Repository) executeCountQuery(ctx context.Context, query string, args []any, logger *slog.Logger) (int, error) {
	var total int
	if err := r.db.QueryRow(ctx, query, args...).Scan(&total); err != nil {
		logger.Error("count query failed", slog.String("query", query), slog.Any("args", args), slog.Any("error", err))
		return 0, fmt.Errorf("%s: %w", "repository.Roles.List", constants.ErrInternal)
	}
	return total, nil
}
