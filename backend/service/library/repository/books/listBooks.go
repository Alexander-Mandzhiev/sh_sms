package books_repository

import (
	"backend/pkg/models/library"
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) ListBooks(ctx context.Context, params *library_models.ListBooksRequest) ([]*library_models.Book, int64, error) {
	const op = "books_repository.ListBooks"
	logger := r.logger.With(slog.String("op", op), slog.String("client_id", params.ClientID.String()), slog.Int("count", int(params.Count)))
	if params.Cursor != nil {
		logger = logger.With(slog.Int64("cursor", *params.Cursor))
	}
	logger.Debug("querying books with pagination")

	query := `SELECT id, client_id, title, author, description, subject_id, class_id, created_at FROM books WHERE client_id = $1 AND deleted_at IS NULL`
	args := []interface{}{params.ClientID}
	argCounter := 2

	if params.Cursor != nil && *params.Cursor > 0 {
		query += fmt.Sprintf(" AND id < $%d", argCounter)
		args = append(args, *params.Cursor)
		argCounter++
	}

	if params.Filter != nil && *params.Filter != "" {
		query += fmt.Sprintf(" AND (title ILIKE $%d OR author ILIKE $%d)", argCounter, argCounter)
		args = append(args, "%"+*params.Filter+"%")
		argCounter++
	}

	query += fmt.Sprintf(" ORDER BY id DESC LIMIT $%d", argCounter)
	args = append(args, params.Count+1)

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		logger.Error("database query failed", slog.Any("error", err))
		return nil, 0, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var books []*library_models.Book
	for rows.Next() {
		var book library_models.Book
		err = rows.Scan(&book.ID, &book.ClientID, &book.Title, &book.Author, &book.Description, &book.SubjectID, &book.ClassID, &book.CreatedAt)
		if err != nil {
			logger.Error("failed to scan book row", slog.Any("error", err))
			return nil, 0, fmt.Errorf("%s: %w", op, err)
		}
		books = append(books, &book)
	}

	if err = rows.Err(); err != nil {
		logger.Error("rows iteration error", slog.Any("error", err))
		return nil, 0, fmt.Errorf("%s: %w", op, err)
	}

	nextCursor := int64(0)
	if len(books) > int(params.Count) {
		nextCursor = books[len(books)-1].ID
		books = books[:len(books)-1]
	}

	logger.Debug("books retrieved", slog.Int("count", len(books)), slog.Bool("has_next", nextCursor > 0), slog.Int64("next_cursor", nextCursor))
	return books, nextCursor, nil
}
