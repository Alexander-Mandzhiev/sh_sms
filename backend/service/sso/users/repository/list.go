package repository

import (
	"backend/service/sso/models"
	"context"
	"log/slog"
)

func (r *Repository) List(ctx context.Context, filter models.ListRequest) ([]models.User, int, error) {
	const op = "repository.User.List"
	logger := r.logger.With(slog.String("op", op), slog.String("client_id", filter.ClientID.String()), slog.Int("page", filter.Page), slog.Int("count", filter.Count))
	logger.Debug("starting users listing")
	return nil, 0, nil
}
