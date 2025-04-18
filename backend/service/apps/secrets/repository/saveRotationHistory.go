package repository

import (
	"backend/service/apps/models"
	"context"
	"github.com/jackc/pgx/v5"
)

func (r *Repository) saveRotationHistory(ctx context.Context, tx pgx.Tx, params models.RotateSecretParams, oldSecret, newSecret string) error {
	const query = `INSERT INTO secret_rotation_history (client_id, app_id, secret_type, old_secret, new_secret, rotated_by) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := tx.Exec(ctx, query, params.ClientID, params.AppID, params.SecretType, oldSecret, newSecret, params.RotatedBy)
	return err
}
