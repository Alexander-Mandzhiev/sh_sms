package repository

import (
	"backend/service/clients/clients/handle"
	"backend/service/clients/clients/models"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	"log/slog"
)

func (r *Repository) Create(ctx context.Context, params *models.CreateParams) (*models.Client, error) {
	const op = "repository.Clients.Create"
	logger := r.logger.With(slog.String("op", op), slog.String("name", params.Name))
	query := `INSERT INTO clients (name, description, type_id, website) VALUES ($1, $2, $3, $4) RETURNING id, is_active, created_at, updated_at`
	var client models.Client
	err := r.db.QueryRow(ctx, query, params.Name, params.Description, params.TypeID, params.Website).Scan(&client.ID, &client.IsActive, &client.CreatedAt, &client.UpdatedAt)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			logger.Warn("duplicate entry", slog.String("constraint", pgErr.ConstraintName))
			return nil, handle.ErrCodeExists
		}
		if errors.As(err, &pgErr) && pgErr.Code == "23503" {
			logger.Warn("invalid foreign key", slog.String("detail", pgErr.Detail))
			return nil, handle.ErrInvalidArgument
		}

		logger.Error("database error", slog.String("error", err.Error()))
		return nil, fmt.Errorf("%s: %w", op, handle.ErrInternal)
	}

	client.Name = params.Name
	client.Description = params.Description
	client.TypeID = params.TypeID
	client.Website = params.Website

	logger.Debug("client created successfully", slog.String("client_id", client.ID))
	return &client, nil
}
