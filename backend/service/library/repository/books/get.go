package books_repository

import (
	library_models "backend/pkg/models/library"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"log/slog"
)

func (r *Repository) GetBookByID(ctx context.Context, id int64, clientID uuid.UUID) (*library_models.Book, error) {
	const op = "repository.Library.Books.GetBookByID"
	logger := r.logger.With(slog.String("op", op), slog.Int64("id", id), slog.String("client_id", clientID.String()))
	logger.Debug("Querying book by ID")

	query := `SELECT id, client_id, title, author, description, subject_id, class_id, created_at FROM books WHERE id = $1 AND client_id = $2 AND deleted_at IS NULL`

	var book library_models.Book
	err := r.db.QueryRow(ctx, query, id, clientID).Scan(&book.ID, &book.ClientID, &book.Title, &book.Author, &book.Description, &book.SubjectID, &book.ClassID, &book.CreatedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logger.Warn("Book not found")
			return nil, library_models.ErrNotFound
		}
		logger.Error("Failed to get book", slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	logger.Debug("Book found", slog.String("title", book.Title))
	return &book, nil
}
