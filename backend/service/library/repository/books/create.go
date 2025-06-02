package books_repository

import (
	"backend/pkg/models/library"
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) CreateBook(ctx context.Context, book *library_models.Book) error {
	const op = "books_repository.CreateBook"
	logger := r.logger.With(slog.String("op", op), slog.String("client_id", book.ClientID.String()), slog.String("title", book.Title))
	logger.Debug("creating book in database")

	query := `INSERT INTO books (client_id, title, author, description, subject_id, class_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, created_at`
	err := r.db.QueryRow(ctx, query, book.ClientID, book.Title, book.Author, book.Description, book.SubjectID, book.ClassID).Scan(&book.ID, &book.CreatedAt)
	if err != nil {
		logger.Error("failed to create book", slog.Any("error", err), slog.Int("subject_id", book.SubjectID), slog.Int("class_id", book.ClassID))
		return fmt.Errorf("%s: %w", op, err)
	}

	logger.Debug("book created in database", slog.Int64("id", book.ID), slog.Time("created_at", book.CreatedAt))
	return nil
}
