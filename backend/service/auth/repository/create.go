package repository

import (
	"backend/service/auth/handle"
	"context"
	"fmt"
	"log/slog"

	"backend/service/auth/models"
)

func (r *Repository) CreateSession(ctx context.Context, session *models.Session) error {
	const op = "repository.Session.CreateSession"
	logger := r.logger.With(slog.String("op", op), slog.String("session_id", session.SessionID.String()), slog.String("user_id", session.UserID.String()))

	query := `INSERT INTO sessions (session_id, user_id, client_id, app_id, access_token_hash, refresh_token_hash, ip_address, user_agent, expires_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	cmdTag, err := r.db.Exec(ctx, query, session.SessionID, session.UserID, session.ClientID, session.AppID, session.AccessTokenHash, session.RefreshTokenHash, session.IPAddress, session.UserAgent, session.ExpiresAt)
	if err != nil {
		logger.Error("failed to execute query", slog.Any("error", err), slog.String("query", query))
		return fmt.Errorf("%s: %w", op, err)
	}

	if cmdTag.RowsAffected() == 0 {
		logger.Warn("no rows affected by insert")
		return fmt.Errorf("%s: %w", op, handle.ErrNoRowsAffected)
	}

	logger.Info("session created successfully", slog.Int64("rows_affected", cmdTag.RowsAffected()))
	return nil
}
