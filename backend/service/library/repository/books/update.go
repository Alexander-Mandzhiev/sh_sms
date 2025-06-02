package books_repository

import (
	"backend/pkg/models/library"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log/slog"
)

func (r *Repository) UpdateBook(ctx context.Context, book *library_models.Book) (*library_models.Book, error) {
	const op = "repository.Library.Books.UpdateBook"
	logger := r.logger.With(slog.String("op", op), slog.Int64("book_id", book.ID))
	logger.Debug("Updating book in repository")

	query := `UPDATE books SET title = $1, author = $2, description = $3, subject_id = $4, class_id = $5
        WHERE id = $6 AND client_id = $7 AND deleted_at IS NULL
        RETURNING id, client_id, title, author, description, subject_id, class_id, created_at`

	var updatedBook library_models.Book
	err := r.db.QueryRow(ctx, query, book.Title, book.Author, book.Description, book.SubjectID, book.ClassID, book.ID, book.ClientID).
		Scan(&updatedBook.ID, &updatedBook.ClientID, &updatedBook.Title, &updatedBook.Author, &updatedBook.Description,
			&updatedBook.SubjectID, &updatedBook.ClassID, &updatedBook.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logger.Warn("role not found for update", slog.String("query", query))
			return nil, library_models.ErrNotFound
		}
		logger.Error("Failed to update book", slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	logger.Info("Book updated successfully", slog.Int64("updated_book_id", updatedBook.ID))
	return &updatedBook, nil
}
