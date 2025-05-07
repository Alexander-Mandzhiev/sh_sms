package repository

import (
	"backend/service/clients/clients/models"
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) List(ctx context.Context, filter models.Filter, pagination models.Pagination) ([]*models.Client, int, error) {
	const op = "repository.Clients.List"
	logger := r.logger.With(slog.String("op", op))
	baseQuery := `SELECT id, name, description, type_id, website, is_active, created_at, updated_at FROM clients WHERE 1=1`
	countQuery := "SELECT COUNT(*) FROM clients WHERE 1=1"
	args := make([]interface{}, 0)
	argNum := 1

	if filter.Search != nil && *filter.Search != "" {
		baseQuery += fmt.Sprintf(" AND name ILIKE '%%' || $%d || '%%'", argNum)
		countQuery += fmt.Sprintf(" AND name ILIKE '%%' || $%d || '%%'", argNum)
		args = append(args, *filter.Search)
		argNum++
	}

	if filter.TypeID != nil {
		baseQuery += fmt.Sprintf(" AND type_id = $%d", argNum)
		countQuery += fmt.Sprintf(" AND type_id = $%d", argNum)
		args = append(args, *filter.TypeID)
		argNum++
	}

	if filter.ActiveOnly != nil {
		baseQuery += fmt.Sprintf(" AND is_active = $%d", argNum)
		countQuery += fmt.Sprintf(" AND is_active = $%d", argNum)
		args = append(args, *filter.ActiveOnly)
		argNum++
	}

	var total int
	err := r.db.QueryRow(ctx, countQuery, args...).Scan(&total)
	if err != nil {
		logger.Error("count query failed", slog.String("error", err.Error()))
		return nil, 0, fmt.Errorf("%s: count failed: %w", op, err)
	}

	baseQuery += fmt.Sprintf(` ORDER BY created_at DESC LIMIT $%d OFFSET $%d`, argNum, argNum+1)
	args = append(args,
		pagination.Count,
		(pagination.Page-1)*pagination.Count,
	)

	rows, err := r.db.Query(ctx, baseQuery, args...)
	if err != nil {
		logger.Error("list query failed", slog.String("error", err.Error()))
		return nil, 0, fmt.Errorf("%s: query failed: %w", op, err)
	}
	defer rows.Close()

	var clients []*models.Client
	for rows.Next() {
		var c models.Client
		err = rows.Scan(&c.ID, &c.Name, &c.Description, &c.TypeID, &c.Website, &c.IsActive, &c.CreatedAt, &c.UpdatedAt)
		if err != nil {
			logger.Error("row scan error", slog.String("error", err.Error()))
			return nil, 0, fmt.Errorf("%s: %w", op, err)
		}
		clients = append(clients, &c)
	}

	logger.Debug("list query executed", slog.Int("total", total), slog.Int("returned", len(clients)))
	return clients, total, nil
}
