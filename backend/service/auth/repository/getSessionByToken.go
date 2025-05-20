package repository

import (
	"backend/service/auth/handle"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"

	"backend/service/auth/models"
	"github.com/jackc/pgx/v5"
)

func (r *Repository) GetSessionByToken(ctx context.Context, tokenHash string) (*models.Session, error) {
	const op = "repository.GetSessionByAccessToken"
	logger := r.logger.With(slog.String("op", op))

	query := `SELECT session_id, user_id, client_id, app_id, access_token_hash, refresh_token_hash, ip_address, user_agent, created_at, last_activity, expires_at, revoked_at
		FROM sessions WHERE (access_token_hash = $1 OR refresh_token_hash = $1) AND revoked_at IS NULL AND expires_at > NOW()`

	var session models.Session
	var revokedAt sql.NullTime

	err := r.db.QueryRow(ctx, query, tokenHash).Scan(&session.SessionID, &session.UserID, &session.ClientID, &session.AppID, &session.AccessTokenHash,
		&session.RefreshTokenHash, &session.IPAddress, &session.UserAgent, &session.CreatedAt, &session.LastActivity, &session.ExpiresAt, &revokedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		logger.Warn("session not found")
		return nil, fmt.Errorf("%w", handle.ErrSessionNotFound)
	}

	if err != nil {
		logger.Error("database error", slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if revokedAt.Valid {
		session.RevokedAt = &revokedAt.Time
	}

	logger.Debug("session found", slog.String("session_id", session.SessionID.String()))
	return &session, nil
}
