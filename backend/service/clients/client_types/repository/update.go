package repository

import (
	"backend/service/clients/client_types/models"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"log/slog"
)

func (r *Repository) Update(ctx context.Context, params *models.UpdateParams) (*models.ClientType, error) {
	const op = "repository.ClientType.Update"
	logger := r.logger.With(slog.String("op", op), slog.Int("id", int(params.ID)))
	logger.Debug("executing update query", slog.String("new_code", params.Code), slog.String("new_name", params.Name))

	query := `UPDATE client_types SET code = $1, name = $2, description = $3, updated_at = NOW() 
                WHERE id = $4 RETURNING id, code, name, description, is_active, created_at, updated_at`
	var ct models.ClientType
	err := r.db.QueryRow(ctx, query, params.Code, params.Name, params.Description, params.ID).Scan(
		&ct.ID, &ct.Code, &ct.Name, &ct.Description, &ct.IsActive, &ct.CreatedAt, &ct.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logger.Warn("client type not found", slog.Int("id", int(params.ID)))
			return nil, fmt.Errorf("%w: %d", ErrNotFound, params.ID)
		}
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			logger.Warn("code conflict", slog.String("code", params.Code), slog.Int("id", int(params.ID)))
			return nil, fmt.Errorf("%w: %s", ErrCodeConflict, params.Code)
		}
		logger.Error("database error", slog.Any("error", err), slog.String("error_type", "database_error"))
		return nil, fmt.Errorf("%w: %v", ErrInternal, err)
	}

	logger.Info("client type updated successfully", slog.Int("id", ct.ID), slog.String("new_code", ct.Code))
	return &ct, nil
}
