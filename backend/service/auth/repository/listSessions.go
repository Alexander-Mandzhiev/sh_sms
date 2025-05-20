package repository

import (
	"backend/service/auth/models"
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"log/slog"
	"strings"
)

func (r *Repository) ListSessions(ctx context.Context, filter models.SessionFilter) ([]models.Session, error) {
	const op = "repository.ListSessions"

	query := strings.Builder{}
	args := make([]interface{}, 0)
	argPos := 1

	query.WriteString(`SELECT session_id, user_id, client_id, app_id, access_token_hash, refresh_token_hash, 
       ip_address, user_agent, created_at, last_activity, expires_at, revoked_at FROM sessions WHERE 1=1`)

	if filter.UserID != uuid.Nil {
		query.WriteString(fmt.Sprintf(" AND user_id = $%d", argPos))
		args = append(args, filter.UserID)
		argPos++
	}

	if filter.ClientID != uuid.Nil {
		query.WriteString(fmt.Sprintf(" AND client_id = $%d", argPos))
		args = append(args, filter.ClientID)
		argPos++
	}

	if filter.AppID > 0 {
		query.WriteString(fmt.Sprintf(" AND app_id = $%d", argPos))
		args = append(args, filter.AppID)
		argPos++
	}

	if filter.ActiveOnly {
		query.WriteString(" AND revoked_at IS NULL AND expires_at > NOW()")
	}

	query.WriteString(" ORDER BY created_at DESC")

	if filter.Count > 0 {
		query.WriteString(fmt.Sprintf(" LIMIT $%d", argPos))
		args = append(args, filter.Count)
		argPos++
	}

	if filter.Page > 0 && filter.Count > 0 {
		offset := (filter.Page - 1) * filter.Count
		query.WriteString(fmt.Sprintf(" OFFSET $%d", argPos))
		args = append(args, offset)
		argPos++
	}

	rows, err := r.db.Query(ctx, query.String(), args...)
	if err != nil {
		r.logger.Error("failed to query sessions", slog.String("op", op), slog.String("user_id", filter.UserID.String()), slog.String("client_id", filter.ClientID.String()), slog.Int("app_id", filter.AppID), slog.Bool("active_only", filter.ActiveOnly), slog.Int("page", filter.Page), slog.Int("count", filter.Count), slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	sessions := make([]models.Session, 0)
	for rows.Next() {
		var session models.Session
		var revokedAt sql.NullTime

		err = rows.Scan(&session.SessionID, &session.UserID, &session.ClientID, &session.AppID, &session.AccessTokenHash,
			&session.RefreshTokenHash, &session.IPAddress, &session.UserAgent, &session.CreatedAt, &session.LastActivity, &session.ExpiresAt, &revokedAt,
		)
		if err != nil {
			r.logger.Error("failed to scan session row", slog.String("op", op), slog.Any("error", err))
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		if revokedAt.Valid {
			session.RevokedAt = &revokedAt.Time
		}
		sessions = append(sessions, session)
	}

	if err = rows.Err(); err != nil {
		r.logger.Error("rows iteration error", slog.String("op", op), slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return sessions, nil
}
