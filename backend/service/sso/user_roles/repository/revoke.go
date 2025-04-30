package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"log/slog"
)

func (r *Repository) Revoke(ctx context.Context, userID uuid.UUID, roleID uuid.UUID, clientID uuid.UUID, appID int) error {
	const op = "repository.UserRoles.Revoke"
	logger := r.logger.With(slog.String("op", op), slog.String("user_id", userID.String()), slog.String("role_id", roleID.String()), slog.String("client_id", clientID.String()), slog.Int("app_id", appID))

	query := `DELETE FROM user_roles WHERE user_id = $1 AND role_id = $2 AND client_id = $3 AND app_id = $4`
	result, err := r.db.Exec(ctx, query, userID, roleID, clientID, appID)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			logger.Error("database error", slog.String("code", pgErr.Code), slog.String("message", pgErr.Message))
			return fmt.Errorf("%s: %w", op, ErrInternal)
		}
		logger.Error("database error", slog.Any("error", err), slog.String("query", query))
		return fmt.Errorf("%s: %w", op, err)
	}

	if result.RowsAffected() == 0 {
		logger.Warn("assignment not found")
		return ErrAssignmentNotFound
	}

	logger.Info("role revoked successfully")
	return nil
}
