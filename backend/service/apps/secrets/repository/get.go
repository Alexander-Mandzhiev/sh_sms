package repository

import (
	"backend/service/apps/constants"
	"backend/service/apps/models"
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5"
)

func (r *Repository) Get(ctx context.Context, clientID string, appID int, secretType string) (*models.Secret, error) {
	const op = "repository.Secret.Get"
	logger := r.logger.With(slog.String("op", op))
	query := `SELECT client_id, app_id, secret_type, current_secret, algorithm, secret_version, generated_at, revoked_at FROM secrets WHERE client_id = $1 AND app_id = $2 AND secret_type = $3`

	var secret models.Secret
	var revokedAt *time.Time
	err := r.db.QueryRow(ctx, query, clientID, appID, secretType).Scan(&secret.ClientID, &secret.AppID, &secret.SecretType, &secret.CurrentSecret, &secret.Algorithm, &secret.SecretVersion, &secret.GeneratedAt, &revokedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logger.Warn("secret not found", slog.String("client_id", clientID), slog.Int("app_id", appID), slog.String("secret_type", secretType))
			return nil, constants.ErrNotFound
		}
		logger.Error("database error", slog.Any("error", err), slog.String("client_id", clientID), slog.Int("app_id", appID), slog.String("secret_type", secretType))
		return nil, fmt.Errorf("database error: %w", err)
	}

	secret.RevokedAt = revokedAt
	logger.Debug("secret retrieved from database", slog.String("client_id", secret.ClientID), slog.Int("app_id", secret.AppID), slog.String("secret_type", secret.SecretType))
	return &secret, nil
}
