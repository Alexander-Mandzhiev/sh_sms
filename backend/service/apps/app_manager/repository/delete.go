package repository

import (
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) Delete(ctx context.Context, id int32) error {
	const op = "repository.Delete"
	logger := r.logger.With(slog.String("op", op))
	result, err := r.db.Exec(ctx, "DELETE FROM apps WHERE id = $1", id)
	if err != nil {
		logger.Error("Failed to delete app", slog.Int("id", int(id)), slog.Any("error", err))
		return fmt.Errorf("%s: %w", op, err)
	}

	if result.RowsAffected() == 0 {
		logger.Warn("Application not found for deletion", slog.Int("id", int(id)))
		return ErrNotFound
	}

	logger.Info("Application deleted", slog.Int("id", int(id)))
	return nil
}
