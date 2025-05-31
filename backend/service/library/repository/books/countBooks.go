package books_repository

import (
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) CountBooks(ctx context.Context, clientID, filter string) (int32, error) {
	const op = "repository.Library.Books.CountBooks"
	logger := r.logger.With(slog.String("op", op))
	logger.Debug("Counting books", slog.String("client_id", clientID), slog.String("filter", filter))

	query := `
		SELECT COUNT(*)
		FROM books
		WHERE client_id = $1
	`
	args := []interface{}{clientID}

	if filter != "" {
		query += ` AND search_vector @@ plainto_tsquery('russian', $2)`
		args = append(args, filter)
	}

	var count int32
	err := r.db.QueryRow(ctx, query, args...).Scan(&count)
	if err != nil {
		logger.Error("Failed to count books", slog.Any("error", err))
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	logger.Debug("Books counted", slog.Int("count", int(count)))
	return count, nil
}
