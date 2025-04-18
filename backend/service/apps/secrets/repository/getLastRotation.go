package repository

import (
	"backend/service/apps/models"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log/slog"
	"time"
)

func (r *Repository) getLastRotation(ctx context.Context, tx pgx.Tx, params models.RotateSecretParams) (time.Time, error) {
	const query = `SELECT rotated_at FROM secret_rotation_history WHERE client_id = $1 AND app_id = $2 AND secret_type = $3 ORDER BY rotated_at DESC LIMIT 1`

	var lastRotation time.Time
	err := tx.QueryRow(ctx, query, params.ClientID, params.AppID, params.SecretType).Scan(&lastRotation)

	if errors.Is(err, pgx.ErrNoRows) {
		return time.Time{}, nil
	}
	if err != nil {
		r.logger.Error("failed to get last rotation", slog.Any("error", err), slog.String("client_id", params.ClientID), slog.Int("app_id", params.AppID), slog.String("secret_type", params.SecretType))
		return time.Time{}, fmt.Errorf("get last rotation failed: %w", err)
	}

	return lastRotation, nil
}
