package repository

import (
	"backend/service/sso/models"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"log/slog"
	"time"
)

func (r *Repository) BatchGetUsers(ctx context.Context, clientID uuid.UUID, userIDs []uuid.UUID) ([]*models.User, error) {
	const op = "repository.BatchGetUsers"
	logger := r.logger.With(slog.String("op", op), slog.String("client_id", clientID.String()), slog.Int("user_count", len(userIDs)))

	if len(userIDs) == 0 {
		logger.Warn("empty user_ids list")
		return nil, fmt.Errorf("%s: %w", op, ErrInvalidArgument)
	}

	ids := make([]string, 0, len(userIDs))
	for _, id := range userIDs {
		ids = append(ids, id.String())
	}

	query := `SELECT id, client_id, email, full_name, phone, is_active, created_at, updated_at, deleted_at FROM users 
        WHERE client_id = $1 AND id = ANY($2) AND deleted_at IS NULL ORDER BY created_at DESC`

	logger.Debug("executing batch get users query", slog.String("query", query), slog.Any("client_id", clientID), slog.Any("user_ids", ids))

	rows, err := r.db.Query(ctx, query, clientID, ids)
	if err != nil {
		logger.Error("failed to execute query", slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", op, ErrInternal)
	}
	defer rows.Close()

	users := make([]*models.User, 0)
	for rows.Next() {
		var user models.User
		var deletedAt *time.Time
		err = rows.Scan(&user.ID, &user.ClientID, &user.Email, &user.FullName, &user.Phone, &user.IsActive, &user.CreatedAt, &user.UpdatedAt, &deletedAt)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				continue
			}
			logger.Error("failed to scan row", slog.Any("error", err))
			return nil, fmt.Errorf("%s: %w", op, ErrInternal)
		}
		users = append(users, &user)
	}

	if err = rows.Err(); err != nil {
		logger.Error("rows iteration error", slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", op, ErrInternal)
	}

	logger.Debug("batch get users completed", slog.Int("found_users", len(users)))
	return users, nil
}
