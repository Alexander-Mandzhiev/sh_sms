package repository

import (
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) CodeExists(ctx context.Context, code string) (bool, error) {
	const op = "repository.ClientType.Exist"
	logger := r.logger.With(slog.String("op", op), slog.String("code", code))
	logger.Debug("processing existing request")
	var exists bool
	const query = `SELECT EXISTS(SELECT 1 FROM client_types WHERE code = $1)`
	err := r.db.QueryRow(ctx, query, code).Scan(&exists)
	if err != nil {
		logger.Warn("failed checking if client type exists: ", err)
		return false, fmt.Errorf("exist check failed: %w", err)
	}

	logger.Debug("client type exists: ", exists)
	return exists, nil
}
