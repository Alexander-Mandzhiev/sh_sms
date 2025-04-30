package repository

import (
	"backend/service/clients/client_types/models"
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5"
)

func (r *Repository) Get(ctx context.Context, id int) (*models.ClientType, error) {
	const op = "repository.ClientType.Get"
	logger := r.logger.With(slog.String("op", op), slog.Int("id", id))
	logger.Debug("processing get request")
	query := `SELECT id, code, name, description, is_active, created_at, updated_at FROM client_types WHERE id = $1`

	var ct models.ClientType
	err := r.db.QueryRow(ctx, query, id).Scan(&ct.ID, &ct.Code, &ct.Name, &ct.Description, &ct.IsActive, &ct.CreatedAt, &ct.UpdatedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logger.Warn("client type not found in DB", slog.Int("id", id))
			return nil, ErrNotFound
		}
		logger.Error("database query failed", slog.String("query", query), slog.Any("error", err))
		return nil, fmt.Errorf("%w: %v", ErrInternal, err)
	}

	logger.Debug("client type successfully retrieved from DB")
	return &ct, nil
}
