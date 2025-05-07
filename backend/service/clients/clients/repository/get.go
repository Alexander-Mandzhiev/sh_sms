package repository

import (
	"backend/service/clients/clients/handle"
	"backend/service/clients/clients/models"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"log/slog"
)

func (r *Repository) GetByID(ctx context.Context, id uuid.UUID) (*models.Client, error) {
	const op = "repository.Clients.GetByID"
	logger := r.logger.With(slog.String("op", op), slog.String("client_id", id.String()))
	query := `SELECT id, name, description, type_id, website, is_active, created_at, updated_at FROM clients WHERE id = $1 AND is_active = TRUE`

	var client models.Client
	err := r.db.QueryRow(ctx, query, id).Scan(&client.ID, &client.Name, &client.Description, &client.TypeID, &client.Website, &client.IsActive, &client.CreatedAt, &client.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logger.Debug("client not found in database")
			return nil, handle.ErrNotFound
		}
		logger.Error("database query failed", slog.String("error", err.Error()), slog.String("query", query))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	logger.Debug("client successfully retrieved")
	return &client, nil
}
