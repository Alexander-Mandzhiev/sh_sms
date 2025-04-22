package repository

import (
	"backend/service/sso/models"
	"context"
	"fmt"
	"log/slog"
	"strings"
)

func (r *Repository) List(ctx context.Context, filter models.ListRequest) ([]models.User, int, error) {
	const op = "repository.User.List"
	logger := r.logger.With(slog.String("op", op), slog.String("client_id", filter.ClientID.String()), slog.Int("page", filter.Page), slog.Int("count", filter.Count))
	logger.Debug("starting users listing")

	whereClauses := []string{"client_id = $1"}
	args := []any{filter.ClientID}

	if filter.EmailFilter != nil && *filter.EmailFilter != "" {
		whereClauses = append(whereClauses, fmt.Sprintf("email ILIKE $%d", len(args)+1))
		args = append(args, "%"+*filter.EmailFilter+"%")
	}

	if filter.PhoneFilter != nil && *filter.PhoneFilter != "" {
		whereClauses = append(whereClauses, fmt.Sprintf("phone LIKE $%d", len(args)+1))
		args = append(args, *filter.PhoneFilter+"%")
	}

	if filter.ActiveOnly != nil && *filter.ActiveOnly {
		whereClauses = append(whereClauses, fmt.Sprintf("is_active = $%d", len(args)+1))
		args = append(args, true)
	}

	whereStr := ""
	if len(whereClauses) > 0 {
		whereStr = "WHERE " + strings.Join(whereClauses, " AND ")
	}

	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM users %s", whereStr)
	var total int
	if err := r.db.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		logger.Error("failed to get total count", slog.Any("error", err))
		return nil, 0, fmt.Errorf("%s: %w", op, err)
	}

	if filter.Count < 1 || filter.Page < 1 {
		return []models.User{}, total, nil
	}

	limit := filter.Count
	offset := (filter.Page - 1) * filter.Count

	query := fmt.Sprintf(`SELECT id, client_id, email, full_name, phone, is_active, created_at, updated_at, deleted_at 
        FROM users %s ORDER BY created_at DESC LIMIT $%d OFFSET $%d`, whereStr, len(args)+1, len(args)+2)
	queryArgs := append(args, limit, offset)
	logger.Debug("executing query", slog.String("query", query), slog.Any("args", queryArgs))

	rows, err := r.db.Query(ctx, query, queryArgs...)
	if err != nil {
		logger.Error("failed to query users", slog.Any("error", err))
		return nil, total, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err = rows.Scan(&user.ID, &user.ClientID, &user.Email, &user.FullName, &user.Phone, &user.IsActive, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt); err != nil {
			logger.Error("failed to scan row", slog.Any("error", err))
			return nil, total, fmt.Errorf("%s: %w", op, err)
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		logger.Error("error during rows iteration", slog.Any("error", err))
		return nil, total, fmt.Errorf("%s: %w", op, err)
	}

	logger.Debug("listing completed", slog.Int("total", total), slog.Int("returned", len(users)))
	return users, total, nil
}
