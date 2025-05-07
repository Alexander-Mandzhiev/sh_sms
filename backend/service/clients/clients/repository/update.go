package repository

import (
	"backend/service/clients/clients/handle"
	"backend/service/clients/clients/models"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"log/slog"
)

func (r *Repository) Update(ctx context.Context, params *models.UpdateParams) (*models.Client, error) {
	const op = "repository.Clients.Update"
	logger := r.logger.With(slog.String("op", op), slog.String("client_id", params.ID))

	if _, err := uuid.Parse(params.ID); err != nil {
		logger.Warn("invalid client ID format", slog.String("error", err.Error()))
		return nil, fmt.Errorf("%w: %s", handle.ErrInvalidArgument, "invalid UUID format")
	}

	query := `UPDATE clients SET name = COALESCE($2, name), description = COALESCE($3, description), type_id = COALESCE($4, type_id), website = COALESCE($5, website), updated_at = NOW()
		WHERE id = $1 RETURNING id, name, description, type_id, website, is_active, created_at, updated_at`

	var client models.Client
	err := r.db.QueryRow(ctx, query, params.ID, params.Name, params.Description, params.TypeID, params.Website).Scan(&client.ID, &client.Name, &client.Description, &client.TypeID, &client.Website, &client.IsActive, &client.CreatedAt, &client.UpdatedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logger.Debug("client not found")
			return nil, handle.ErrNotFound
		}
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			logger.Warn("duplicate entry", slog.String("constraint", pgErr.ConstraintName), slog.String("detail", pgErr.Detail))
			return nil, handle.ErrCodeExists
		}
		if errors.As(err, &pgErr) && pgErr.Code == "23503" {
			logger.Warn("foreign key violation", slog.String("detail", pgErr.Detail))
			return nil, handle.ErrInvalidArgument
		}

		logger.Error("database error", slog.String("error", err.Error()))
		return nil, fmt.Errorf("%s: %w", op, handle.ErrInternal)
	}

	logger.Debug("client updated successfully", slog.String("updated_at", client.UpdatedAt.String()))
	return &client, nil
}
