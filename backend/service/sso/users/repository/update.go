package repository

import (
	"backend/service/constants"
	"backend/service/sso/models"
	"context"
	"fmt"
	"github.com/google/uuid"
	"log/slog"
	"strings"
	"time"
)

func (r *Repository) Update(ctx context.Context, userID, clientID uuid.UUID, update models.UserUpdate) error {
	const op = "repository.User.Update"
	logger := r.logger.With(slog.String("op", op), slog.String("user_id", userID.String()), slog.String("client_id", clientID.String()))
	logger.Debug("attempting partial user update")

	query := "UPDATE users SET"
	args := make([]interface{}, 0)
	params := make([]string, 0)
	fields := make([]string, 0)

	if update.Email != nil {
		params = append(params, fmt.Sprintf("email = $%d", len(args)+1))
		args = append(args, *update.Email)
		fields = append(fields, "email")
	}
	if update.FullName != nil {
		params = append(params, fmt.Sprintf("full_name = $%d", len(args)+1))
		args = append(args, *update.FullName)
		fields = append(fields, "full_name")
	}
	if update.Phone != nil {
		params = append(params, fmt.Sprintf("phone = $%d", len(args)+1))
		args = append(args, *update.Phone)
		fields = append(fields, "phone")
	}
	if update.IsActive != nil {
		params = append(params, fmt.Sprintf("is_active = $%d", len(args)+1))
		args = append(args, *update.IsActive)
		fields = append(fields, "is_active")
	}

	if len(params) == 0 {
		logger.Warn("no fields to update")
		return fmt.Errorf("%s: %w", op, constants.ErrNoFieldsToUpdate)
	}

	params = append(params, fmt.Sprintf("updated_at = $%d", len(args)+1))
	args = append(args, time.Now().UTC())
	fields = append(fields, "updated_at")

	query += " " + strings.Join(params, ", ") + fmt.Sprintf(" WHERE id = $%d AND client_id = $%d", len(args)+1, len(args)+2)
	args = append(args, userID, clientID)

	logger = logger.With(slog.String("query", query), slog.Any("fields", fields))

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		logger.Error("database operation failed", slog.Any("error", err), slog.Any("args", args))
		return fmt.Errorf("%s: %w", op, err)
	}

	if rows := result.RowsAffected(); rows == 0 {
		logger.Warn("no rows affected - user not found or data not changed")
		return fmt.Errorf("%s: %w", op, constants.ErrNotFound)
	}

	logger.Info("user successfully updated", slog.Int64("rows_affected", result.RowsAffected()))
	return nil
}
