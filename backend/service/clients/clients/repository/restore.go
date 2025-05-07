package repository

import (
	"backend/service/clients/clients/handle"
	"backend/service/clients/clients/models"
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"log/slog"
)

func (r *Repository) Restore(ctx context.Context, id uuid.UUID) (*models.Client, error) {
	const op = "repository.Restore"
	logger := r.logger.With(slog.String("op", op), slog.String("client_id", id.String()))
	query := `UPDATE clients SET is_active = TRUE, updated_at = NOW() WHERE id = $1 AND is_active = FALSE
		RETURNING id, name, description, type_id, website, is_active, created_at, updated_at`

	var client models.Client
	err := r.db.QueryRow(ctx, query, id).Scan(&client.ID, &client.Name, &client.Description, &client.TypeID, &client.Website, &client.IsActive, &client.CreatedAt, &client.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logger.Debug("client not found or already active")
			return nil, handle.ErrNotFound
		}
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			logger.Warn("duplicate entry after restore", slog.String("constraint", pgErr.ConstraintName))
			return nil, handle.ErrCodeExists
		}
		if errors.As(err, &pgErr) && pgErr.Code == "23503" {
			logger.Warn("invalid type_id after restore", slog.String("detail", pgErr.Detail))
			return nil, handle.ErrInvalidArgument
		}
		logger.Error("database error", slog.String("error", err.Error()))
		return nil, handle.ErrInternal
	}

	logger.Debug("client restored successfully")
	return &client, nil
}
