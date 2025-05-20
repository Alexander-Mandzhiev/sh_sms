package repository

import (
	"backend/service/auth/handle"
	"backend/service/auth/models"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log/slog"
	"time"
)

func (r *Repository) GetSessionByTokenHash(ctx context.Context, accessHash string, refreshHash string) (*models.Session, error) {
	const op = "repository.Session.GetSessionByHashToken"
	logger := r.logger.With(slog.String("op", op), slog.String("access_hash", maskHash(accessHash)), slog.String("refresh_hash", maskHash(refreshHash)))

	query := `SELECT 
        session_id, user_id, client_id, app_id, 
        access_token_hash, refresh_token_hash, 
        ip_address, user_agent, created_at, 
        last_activity, expires_at, revoked_at 
    FROM sessions 
    WHERE 
        access_token_hash = $1 
        AND refresh_token_hash = $2 
        AND revoked_at IS NULL 
        AND expires_at > $3`

	var session models.Session
	err := r.db.QueryRow(ctx, query, accessHash, refreshHash, time.Now()).Scan(
		&session.SessionID, &session.UserID, &session.ClientID, &session.AppID, &session.AccessTokenHash, &session.RefreshTokenHash,
		&session.IPAddress, &session.UserAgent, &session.CreatedAt, &session.LastActivity, &session.ExpiresAt, &session.RevokedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logger.Warn("session not found")
			return nil, fmt.Errorf("%s: %w", op, handle.ErrSessionNotFound)
		}
		logger.Error("database error", slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	logger.Debug("session found", slog.String("session_id", session.SessionID.String()))
	return &session, nil
}
