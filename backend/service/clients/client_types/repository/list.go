package repository

import (
	"backend/service/clients/client_types/models"
	"context"
	"fmt"
	"log/slog"
	"strings"
)

func (r *Repository) List(ctx context.Context, filter models.Filter, pagination models.Pagination) ([]*models.ClientType, int, error) {
	const op = "repository.ClientType.List"
	logger := r.logger.With(slog.String("op", op), slog.Any("filter", filter), slog.Any("pagination", pagination))

	baseQuery := "SELECT id, code, name, description, is_active, created_at, updated_at FROM client_types"
	countQuery := "SELECT COUNT(*) FROM client_types"

	var conditions []string
	var args []interface{}
	argPos := 1

	if filter.Search != nil && *filter.Search != "" {
		conditions = append(conditions, fmt.Sprintf("(code ILIKE $%d OR name ILIKE $%d)", argPos, argPos))
		args = append(args, "%"+*filter.Search+"%")
		argPos++
	}

	if filter.ActiveOnly != nil {
		conditions = append(conditions, fmt.Sprintf("is_active = $%d", argPos))
		args = append(args, *filter.ActiveOnly)
		argPos++
	}

	whereClause := ""
	if len(conditions) > 0 {
		whereClause = " WHERE " + strings.Join(conditions, " AND ")
	}

	var total int
	err := r.db.QueryRow(ctx, countQuery+whereClause, args...).Scan(&total)
	if err != nil {
		logger.Error("failed to get total count", slog.Any("error", err))
		return nil, 0, fmt.Errorf("%w: count failed", ErrInternal)
	}

	query := baseQuery + whereClause + " ORDER BY created_at DESC" + fmt.Sprintf(" LIMIT $%d OFFSET $%d", argPos, argPos+1)

	args = append(args, pagination.PageSize, (pagination.Page-1)*pagination.PageSize)

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		logger.Error("query failed", slog.String("query", query), slog.Any("args", args), slog.Any("error", err))
		return nil, 0, fmt.Errorf("%w: query failed", ErrInternal)
	}
	defer rows.Close()

	var result []*models.ClientType
	for rows.Next() {
		var ct models.ClientType
		if err = rows.Scan(&ct.ID, &ct.Code, &ct.Name, &ct.Description, &ct.IsActive, &ct.CreatedAt, &ct.UpdatedAt); err != nil {
			logger.Error("row scan failed", slog.Any("error", err))
			return nil, 0, fmt.Errorf("%w: scan failed", ErrInternal)
		}
		result = append(result, &ct)
	}

	if err = rows.Err(); err != nil {
		logger.Error("rows iteration error", slog.Any("error", err))
		return nil, 0, fmt.Errorf("%w: rows error", ErrInternal)
	}

	logger.Debug("successfully retrieved records", slog.Int("count", len(result)), slog.Int("total", total))
	return result, total, nil
}
