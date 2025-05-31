package classes_repository

import (
	"backend/pkg/models/library"
	classes_handle "backend/service/library/handle/classes"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log/slog"
)

func (r *Repository) GetClassByID(ctx context.Context, id int) (*library_models.Class, error) {
	const op = "repository.Library.Classes.GetClassByID"
	logger := r.logger.With(slog.String("op", op), slog.Int("id", id))
	logger.Debug("querying class by ID")

	query := `SELECT id, grade FROM classes WHERE id = $1`

	var class library_models.Class
	err := r.db.QueryRow(ctx, query, id).Scan(&class.ID, &class.Grade)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logger.Warn("class not found")
			return nil, classes_handle.ErrClassNotFound
		}

		logger.Error("failed to get class", slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	logger.Debug("class found", slog.Int("grade", int(class.Grade)))
	return &class, nil
}
