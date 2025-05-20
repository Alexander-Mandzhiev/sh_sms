package repository

import (
	"backend/service/auth/handle"
	"context"
	"fmt"

	"github.com/google/uuid"
	"log/slog"
)

func (r *Repository) DeleteSession(ctx context.Context, sessionID uuid.UUID) error {
	const op = "repository.DeleteSession"
	query := `DELETE FROM sessions WHERE session_id = $1`

	result, err := r.db.Exec(ctx, query, sessionID)
	if err != nil {
		r.logger.Error("failed to delete session", slog.String("op", op), slog.String("session_id", sessionID.String()), slog.Any("error", err))
		return fmt.Errorf("%s: %w", op, err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		r.logger.Warn("session not found", slog.String("op", op), slog.String("session_id", sessionID.String()))
		return handle.ErrSessionNotFound
	}

	r.logger.Debug("session deleted successfully", slog.String("op", op), slog.String("session_id", sessionID.String()))

	return nil
}
