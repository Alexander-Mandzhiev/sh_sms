package repository

import (
	"backend/service/apps/constants"
	"backend/service/apps/models"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	"log/slog"
	"time"
)

func (r *Repository) Create(ctx context.Context, params models.CreateSecretParams) (*models.Secret, error) {
	const op = "repository.Secret.CreateSecret"
	logger := r.logger.With(slog.String("op", op))
	secretValue, err := generateSecureSecret()
	if err != nil {
		logger.Error("failed to generate secret", slog.Any("error", err))
		return nil, fmt.Errorf("%w: %v", constants.ErrInternal, err)
	}

	newSecret := &models.Secret{
		ClientID:      params.ClientID,
		AppID:         params.AppID,
		SecretType:    params.SecretType,
		CurrentSecret: secretValue,
		Algorithm:     params.Algorithm,
		SecretVersion: 1,
		GeneratedAt:   time.Now(),
	}

	query := `INSERT INTO secrets (client_id, app_id, secret_type, current_secret, algorithm, secret_version, generated_at, revoked_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err = r.db.Exec(ctx, query, newSecret.ClientID, newSecret.AppID, newSecret.SecretType, newSecret.CurrentSecret, newSecret.Algorithm, newSecret.SecretVersion, newSecret.GeneratedAt, newSecret.RevokedAt)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			logger.Warn("secret already exists")
			return nil, constants.ErrSecretAlreadyExists
		}
		logger.Error("failed to create secret", slog.Any("error", err))
		return nil, fmt.Errorf("create secret failed: %w", err)
	}

	logger.Debug("secret created successfully")
	return newSecret, nil
}
