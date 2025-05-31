package classes_repository

import (
	"backend/pkg/models/library"
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) ListClasses(ctx context.Context) ([]*library_models.Class, error) {
	const op = "repository.Library.Classes.ListClasses"
	logger := r.logger.With(slog.String("op", op))
	logger.Debug("querying all classes")

	query := `SELECT id, grade FROM classes ORDER BY grade ASC`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		logger.Error("failed to query classes", slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	classes := make([]*library_models.Class, 0, 11)
	for rows.Next() {
		var class library_models.Class
		if err = rows.Scan(&class.ID, &class.Grade); err != nil {
			logger.Error("failed to scan class row", slog.Any("error", err))
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		classes = append(classes, &class)
	}

	if err = rows.Err(); err != nil {
		logger.Error("error during rows iteration", slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	logger.Debug("classes retrieved", slog.Int("count", len(classes)))
	return classes, nil
}
