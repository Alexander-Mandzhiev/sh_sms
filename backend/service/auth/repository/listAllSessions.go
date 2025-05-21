package repository

import (
	"backend/service/auth/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"log/slog"
	"strings"
)

func (r *Repository) ListAllSessions(ctx context.Context, filter models.AllSessionsFilter) ([]models.Session, error) {
	const op = "repository.ListAllSessions"
	logger := r.logger.With(slog.String("op", op), slog.String("client_id", filter.ClientID.String()), slog.Int("app_id", filter.AppID))

	query := strings.Builder{}
	args := make([]interface{}, 0)
	argPos := 1

	query.WriteString(`SELECT session_id, user_id, client_id, app_id, access_token_hash, refresh_token_hash,
            ip_address, user_agent, created_at, last_activity, expires_at, revoked_at FROM sessions WHERE 1=1`)

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

	if filter.ActiveOnly != nil && *filter.ActiveOnly {
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

	logger.Debug("executing query", slog.String("query", query.String()), slog.Any("args", args))

	rows, err := r.db.Query(ctx, query.String(), args...)
	if err != nil {
		logger.Error("query execution failed", slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	sessions := make([]models.Session, 0)
	for rows.Next() {
		var session models.Session
		var revokedAt sql.NullTime

		err = rows.Scan(&session.SessionID, &session.UserID, &session.ClientID, &session.AppID, &session.AccessTokenHash,
			&session.RefreshTokenHash, &session.IPAddress, &session.UserAgent, &session.CreatedAt, &session.LastActivity, &session.ExpiresAt, &revokedAt)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				continue
			}
			logger.Error("row scan error", slog.Any("error", err))
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		if revokedAt.Valid {
			session.RevokedAt = &revokedAt.Time
		}

		sessions = append(sessions, session)
	}

	if err = rows.Err(); err != nil {
		logger.Error("rows iteration error", slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	logger.Debug("sessions fetched", slog.Int("count", len(sessions)), slog.Bool("has_results", len(sessions) > 0))
	return sessions, nil
}
