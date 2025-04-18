package repository

import (
	"backend/service/apps/models"
	"backend/service/constants"
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5"
)

func (r *Repository) Revoke(ctx context.Context, clientID string, appID int, secretType string) (*models.Secret, error) {
	const op = "repository.Secret.Revoke"
	logger := r.logger.With(slog.String("op", op))

	tx, err := r.db.Begin(ctx)
	if err != nil {
		logger.Error("failed to begin transaction", slog.Any("error", err))
		return nil, fmt.Errorf("transaction start failed: %w", err)
	}
	defer tx.Rollback(ctx)

	query := `SELECT client_id, app_id, secret_type, current_secret, algorithm, secret_version, generated_at, revoked_at FROM secrets WHERE client_id = $1 AND app_id = $2 AND secret_type = $3 FOR UPDATE`
	var secret models.Secret
	err = tx.QueryRow(ctx, query,
		clientID, appID, secretType).Scan(&secret.ClientID, &secret.AppID, &secret.SecretType, &secret.CurrentSecret, &secret.Algorithm, &secret.SecretVersion, &secret.GeneratedAt, &secret.RevokedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logger.Warn("secret not found")
			return nil, constants.ErrNotFound
		}
		logger.Error("failed to get secret", slog.Any("error", err))
		return nil, fmt.Errorf("get secret failed: %w", err)
	}

	if secret.RevokedAt != nil {
		logger.Warn("secret already revoked")
		return nil, constants.ErrAlreadyExists
	}

	query = `UPDATE secrets SET revoked_at = $1 WHERE client_id = $2 AND app_id = $3 AND secret_type = $4`
	now := time.Now().UTC()
	_, err = tx.Exec(ctx, query, now, clientID, appID, secretType)

	if err != nil {
		logger.Error("failed to revoke secret", slog.Any("error", err))
		return nil, fmt.Errorf("revoke secret failed: %w", err)
	}

	if err = tx.Commit(ctx); err != nil {
		logger.Error("transaction commit failed", slog.Any("error", err))
		return nil, fmt.Errorf("transaction commit failed: %w", err)
	}

	secret.RevokedAt = &now
	logger.Debug("secret revoked successfully")
	return &secret, nil
}
