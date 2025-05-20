package repository

import (
	"backend/service/auth/handle"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"log/slog"
)

func (r *Repository) UpdateTokens(ctx context.Context, sessionID uuid.UUID, accessTokenHash string, refreshTokenHash string, expiresAt time.Time) error {
	const op = "repository.Session.UpdateTokens"
	logger := r.logger.With(slog.String("op", op), slog.String("session_id", sessionID.String()))

	query := `UPDATE sessions SET access_token_hash = $1, refresh_token_hash = $2, expires_at = $3, last_activity = NOW() WHERE session_id = $4`

	tag, err := r.db.Exec(ctx, query, accessTokenHash, refreshTokenHash, expiresAt, sessionID)
	if err != nil {
		logger.Error("failed to execute query", slog.Any("error", err), slog.String("query", query))
		return fmt.Errorf("%s: %w", op, err)
	}

	if tag.RowsAffected() == 0 {
		logger.Warn("no session found for update")
		return fmt.Errorf("%s: %w", op, handle.ErrSessionNotFound)
	}

	logger.Debug("tokens updated successfully", slog.Int64("rows_affected", tag.RowsAffected()))
	return nil
}
