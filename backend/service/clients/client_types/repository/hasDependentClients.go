package repository

import (
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) HasDependentClients(ctx context.Context, typeID int) (bool, error) {
	const op = "repository.ClientType.HasDependentClients"
	logger := r.logger.With(slog.String("op", op), slog.Int("type_id", typeID))
	logger.Debug("checking dependent clients")
	const query = `SELECT EXISTS(SELECT 1 FROM clients WHERE type_id = $1)`
	var exists bool
	err := r.db.QueryRow(ctx, query, typeID).Scan(&exists)
	if err != nil {
		logger.Error("dependency check failed", slog.String("query", query), slog.Any("error", err))
		return false, fmt.Errorf("%s: %w", op, err)
	}

	logger.Debug("dependency check result", slog.Bool("exists", exists))
	return exists, nil
}
