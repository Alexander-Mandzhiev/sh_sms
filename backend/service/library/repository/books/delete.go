package books_repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"log/slog"
)

func (r *Repository) DeleteBook(ctx context.Context, id int64, clientID uuid.UUID) error {
	const op = "repository.Library.Books.DeleteBook"
	logger := r.logger.With(slog.String("op", op), slog.Int64("id", id), slog.String("client_id", clientID.String()))
	logger.Debug("Soft deleting book by ID")

	query := `UPDATE books
        SET deleted_at = NOW()
        WHERE id = $1 AND client_id = $2 AND deleted_at IS NULL
        RETURNING id`

	var deletedID int64
	err := r.db.QueryRow(ctx, query, id, clientID).Scan(&deletedID)
	if err != nil {
		logger.Error("Failed to soft delete book", "error", err)
		return fmt.Errorf("%s: %w", op, err)
	}

	logger.Info("Book successfully soft deleted", slog.Int64("deleted_id", deletedID))
	return nil
}
