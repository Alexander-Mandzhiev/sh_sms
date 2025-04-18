package repository

import (
	"backend/service/apps/constants"
	"backend/service/apps/models"
	"backend/service/utils"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log/slog"
)

func (r *Repository) GetRotation(ctx context.Context, id int) (*models.RotationHistory, error) {
	const op = "repository.Secret.GetRotation"
	logger := r.logger.With(slog.String("op", op), slog.Int("id", id))
	logger.Debug("processing rotation history request")

	query := `SELECT id, client_id, app_id, secret_type, old_secret, new_secret, rotated_by, rotated_at FROM secret_rotation_history WHERE id = $1`

	var history models.RotationHistory

	err := r.db.QueryRow(ctx, query, id).Scan(&history.ID, &history.ClientID, &history.AppID,
		&history.SecretType, &history.OldSecret, &history.NewSecret, &history.RotatedBy, &history.RotatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logger.Warn("rotation history not found", slog.Int("id", id))
			return nil, constants.ErrNotFound
		}
		logger.Error("database operation failed", slog.Any("error", err), slog.Int("id", id))
		return nil, fmt.Errorf("%w: database error", constants.ErrInternal)
	}

	if history.ClientID == "" || history.AppID <= 0 || !utils.IsValidSecretType(history.SecretType) {
		logger.Error("invalid history data", slog.String("client_id", history.ClientID), slog.Int("app_id", history.AppID), slog.String("secret_type", history.SecretType))
		return nil, constants.ErrInternal
	}

	logger.Info("rotation history retrieved")
	return &history, nil
}
