package repository

import (
	"backend/service/auth/handle"
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/google/uuid"
)

func (r *Repository) RevokeSession(ctx context.Context, sessionID uuid.UUID) error {
	const op = "repository.Session.RevokeSession"
	logger := r.logger.With(slog.String("op", op), slog.String("session_id", sessionID.String()))

	query := `UPDATE sessions SET revoked_at = $1 WHERE session_id = $2 AND revoked_at IS NULL`

	result, err := r.db.Exec(ctx, query, time.Now().UTC(), sessionID)
	if err != nil {
		logger.Error("update failed", slog.Any("error", err))
		return fmt.Errorf("%s: %w", op, err)
	}

	if result.RowsAffected() == 0 {
		logger.Warn("no session revoked")
		return fmt.Errorf("%s: %w", op, handle.ErrSessionNotFound)
	}

	logger.Info("session revoked successfully")
	return nil
}

func maskHash(hash string) string {
	if len(hash) < 8 {
		return "***"
	}
	return hash[:2] + "***" + hash[len(hash)-2:]
}
