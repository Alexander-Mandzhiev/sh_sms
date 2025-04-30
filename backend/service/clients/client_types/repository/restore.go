package repository

import (
	"backend/service/clients/client_types/models"
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *Repository) Restore(ctx context.Context, id int) (*models.ClientType, error) {
	const op = "repository.ClientType.Restore"
	logger := r.logger.With(slog.String("op", op), slog.Int("id", id))
	logger.Debug("attempting to restore client type")

	query := `UPDATE client_types SET is_active = true, updated_at = NOW() WHERE id = $1 AND is_active = false 
				RETURNING id, code, name, description, is_active, created_at, updated_at`

	var ct models.ClientType
	err := r.db.QueryRow(ctx, query, id).Scan(&ct.ID, &ct.Code, &ct.Name, &ct.Description, &ct.IsActive, &ct.CreatedAt, &ct.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			exists, existErr := r.Exist(ctx, id)
			if existErr != nil {
				logger.Error("existence check failed", slog.Any("error", existErr))
				return nil, fmt.Errorf("%w: existence check failed", ErrInternal)
			}
			if exists {
				logger.Warn("client type already active")
				return nil, ErrAlreadyActive
			}
			logger.Warn("client type not found")
			return nil, ErrNotFound
		}

		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			logger.Warn("code conflict during restore", slog.String("code", pgErr.Code), slog.String("detail", pgErr.Detail))
			return nil, fmt.Errorf("%w: %v", ErrCodeConflict, pgErr.Detail)
		}

		logger.Error("restore operation failed", slog.Any("error", err))
		return nil, fmt.Errorf("%w: restore failed", ErrInternal)
	}

	logger.Info("client type restored successfully", slog.String("code", ct.Code), slog.Time("updated_at", ct.UpdatedAt))
	return &ct, nil
}
