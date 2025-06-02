package books_repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"log/slog"
)

func (r *Repository) CountBooks(ctx context.Context, clientID uuid.UUID, filter *string) (int32, error) {
	const op = "books_repository.CountBooks"
	logger := r.logger.With(slog.String("op", op), slog.String("client_id", clientID.String()))
	logger.Debug("counting books")

	query := `SELECT COUNT(*) FROM books WHERE client_id = $1 AND deleted_at IS NULL`
	args := []interface{}{clientID}
	argCounter := 2

	if filter != nil && *filter != "" {
		query += fmt.Sprintf(" AND (title ILIKE $%d OR author ILIKE $%d)", argCounter, argCounter)
		args = append(args, "%"+*filter+"%")
		argCounter++
	}

	var count int32
	err := r.db.QueryRow(ctx, query, args...).Scan(&count)
	if err != nil {
		logger.Error("failed to count books", slog.Any("error", err))
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	logger.Debug("books counted", slog.Int("count", int(count)))
	return count, nil
}
