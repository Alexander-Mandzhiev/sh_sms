package repository

import (
	"backend/service/apps/constants"
	"backend/service/apps/models"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log/slog"
	"time"
)

func (r *Repository) Rotate(ctx context.Context, params models.RotateSecretParams) (*models.Secret, error) {
	const op = "repository.Secret.Rotate"
	logger := r.logger.With(slog.String("op", op))

	tx, err := r.db.Begin(ctx)
	if err != nil {
		logger.Error("failed to begin transaction", slog.Any("error", err))
		return nil, fmt.Errorf("transaction start failed: %w", err)
	}
	defer tx.Rollback(ctx)

	query := `SELECT client_id, app_id, secret_type, current_secret, algorithm, secret_version, generated_at, revoked_at FROM secrets WHERE client_id = $1 AND app_id = $2 AND secret_type = $3 FOR UPDATE`
	var oldSecret models.Secret
	err = tx.QueryRow(ctx, query, params.ClientID, params.AppID, params.SecretType).Scan(&oldSecret.ClientID, &oldSecret.AppID, &oldSecret.SecretType, &oldSecret.CurrentSecret, &oldSecret.Algorithm, &oldSecret.SecretVersion, &oldSecret.GeneratedAt, &oldSecret.RevokedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logger.Warn("secret not found")
			return nil, constants.ErrNotFound
		}
		logger.Error("failed to get secret", slog.Any("error", err))
		return nil, fmt.Errorf("get secret failed: %w", err)
	}

	newSecretValue, err := generateSecureSecret()
	if err != nil {
		logger.Error("failed to generate secret", slog.Any("error", err))
		return nil, fmt.Errorf("secret generation failed: %w", err)
	}

	newSecret := &models.Secret{
		ClientID:      oldSecret.ClientID,
		AppID:         oldSecret.AppID,
		SecretType:    oldSecret.SecretType,
		CurrentSecret: newSecretValue,
		Algorithm:     oldSecret.Algorithm,
		SecretVersion: oldSecret.SecretVersion + 1,
		GeneratedAt:   time.Now(),
	}
	query = `UPDATE secrets SET current_secret = $1, secret_version = $2, generated_at = $3, revoked_at = NULL WHERE client_id = $4 AND app_id = $5 AND secret_type = $6`
	_, err = tx.Exec(ctx, query, newSecret.CurrentSecret, newSecret.SecretVersion, newSecret.GeneratedAt, newSecret.ClientID, newSecret.AppID, newSecret.SecretType)

	if err != nil {
		logger.Error("failed to update secret", slog.Any("error", err))
		return nil, fmt.Errorf("update secret failed: %w", err)
	}

	query = `INSERT INTO secret_rotation_history (client_id, app_id, secret_type, old_secret, new_secret, rotated_by) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err = tx.Exec(ctx, query, params.ClientID, params.AppID, params.SecretType, oldSecret.CurrentSecret, newSecret.CurrentSecret, params.RotatedBy)

	if err != nil {
		logger.Error("failed to save rotation history", slog.Any("error", err))
		return nil, fmt.Errorf("save history failed: %w", err)
	}

	if err = tx.Commit(ctx); err != nil {
		logger.Error("transaction commit failed", slog.Any("error", err))
		return nil, fmt.Errorf("transaction commit failed: %w", err)
	}

	logger.Debug("secret rotated successfully")
	return newSecret, nil
}
