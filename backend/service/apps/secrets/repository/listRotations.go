package repository

import (
	"backend/service/apps/models"
	"context"
	"log/slog"
)

func (r *Repository) ListRotations(ctx context.Context, filter models.ListFilter) ([]*models.RotationHistory, int, error) {
	const op = "repository.Secret.ListRotations"
	logger := r.logger.With(slog.String("op", op))

	return nil, 0, nil
}
