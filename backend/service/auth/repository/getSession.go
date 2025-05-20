package repository

import (
	"backend/service/auth/handle"
	"backend/service/auth/models"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"log/slog"
)

// AND expires_at > NOW()
func (r *Repository) GetSession(ctx context.Context, userID uuid.UUID, clientID uuid.UUID, appID int, refreshTokenHash string) (*models.Session, error) {
	const op = "repository.Session.GetSession"
	logger := r.logger.With(slog.String("op", op), slog.String("user_id", userID.String()), slog.String("client_id", clientID.String()), slog.Int("app_id", appID))

	query := `SELECT 
        session_id, user_id, client_id, app_id, 
        access_token_hash, refresh_token_hash, 
        ip_address, user_agent, created_at, 
        last_activity, expires_at, revoked_at 
    FROM sessions 
    WHERE 
        user_id = $1 
        AND client_id = $2 
        AND app_id = $3 
        AND refresh_token_hash = $4 
        AND revoked_at IS NULL`

	var session models.Session

	err := r.db.QueryRow(ctx, query, userID, clientID, appID, refreshTokenHash).Scan(
		&session.SessionID,
		&session.UserID,
		&session.ClientID,
		&session.AppID,
		&session.AccessTokenHash,
		&session.RefreshTokenHash,
		&session.IPAddress,
		&session.UserAgent,
		&session.CreatedAt,
		&session.LastActivity,
		&session.ExpiresAt,
		&session.RevokedAt,
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
