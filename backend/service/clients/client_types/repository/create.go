package repository

import (
	"backend/service/clients/client_types/models"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	"log/slog"
)

func (r *Repository) Create(ctx context.Context, params *models.CreateParams) (*models.ClientType, error) {
	const op = "repository.ClientType.Create"
	logger := r.logger.With(slog.String("op", op), slog.String("code", params.Code), slog.String("name", params.Name), slog.Bool("is_active", *params.IsActive))
	logger.Debug("attempting to create client type")
	query := `INSERT INTO client_types (code, name, description, is_active) VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at`
	var ct models.ClientType
	err := r.db.QueryRow(ctx, query, params.Code, params.Name, safeString(params.Description), *params.IsActive).Scan(&ct.ID, &ct.CreatedAt, &ct.UpdatedAt)

	if err != nil {
		logger.Error("database operation failed", slog.String("error", err.Error()))
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return nil, fmt.Errorf("%w: %s", ErrCodeExists, params.Code)
		}
		return nil, fmt.Errorf("%w: %v", ErrInternal, err)
	}

	ct.Code = params.Code
	ct.Name = params.Name
	ct.Description = safeString(params.Description)
	ct.IsActive = *params.IsActive

	logger.Info("client type created", slog.Int("id", ct.ID))
	return &ct, nil
}

func safeString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
