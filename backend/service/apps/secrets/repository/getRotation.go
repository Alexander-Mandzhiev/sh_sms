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

func (r *Repository) GetRotation(ctx context.Context, clientID string, appID int, secretType string, rotatedAt time.Time) (*models.RotationHistory, error) {
	const op = "repository.Secret.GetRotation"
	logger := r.logger.With(slog.String("op", op), slog.String("client_id", clientID), slog.Int("app_id", appID), slog.String("secret_type", secretType), slog.Time("rotated_at", rotatedAt))
	query := `SELECT client_id, app_id, secret_type, old_secret, new_secret, rotated_by, rotated_at FROM secret_rotation_history WHERE client_id = $1 AND app_id = $2 AND secret_type = $3 AND rotated_at = $4`
	var history models.RotationHistory
	var rotatedBy *string

	err := r.db.QueryRow(ctx, query, clientID, appID, secretType, rotatedAt).Scan(
		&history.ClientID, &history.AppID, &history.SecretType, &history.OldSecret, &history.NewSecret, &rotatedBy, &history.RotatedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logger.Warn("rotation history not found")
			return nil, constants.ErrNotFound
		}
		logger.Error("database error", slog.Any("error", err))
		return nil, fmt.Errorf("%w: database error", constants.ErrInternal)
	}

	if rotatedBy != nil {
		history.RotatedBy = *rotatedBy
	}

	logger.Debug("rotation history retrieved successfully")
	return &history, nil
}
