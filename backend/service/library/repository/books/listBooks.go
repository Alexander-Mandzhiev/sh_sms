package books_repository

import (
	library_models "backend/pkg/models/library"
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) ListBooks(ctx context.Context, params *library_models.ListBooksParams) ([]*library_models.Book, error) {
	const op = "repository.Library.Books.ListBooks"
	logger := r.logger.With(slog.String("op", op))
	logger.Debug("Querying books", slog.String("client_id", params.ClientID), slog.Int("page_size", int(params.PageSize)), slog.String("filter", params.Filter))

	query := `
		SELECT id, client_id, title, author, description, subject_id, class_id, created_at
		FROM books
		WHERE client_id = $1
	`

	args := []interface{}{params.ClientID}
	argCount := 1

	if params.Filter != "" {
		argCount++
		query += fmt.Sprintf(` AND search_vector @@ plainto_tsquery('russian', $%d)`, argCount)
		args = append(args, params.Filter)
	}

	if params.Cursor != nil {
		argCount++
		query += fmt.Sprintf(` 
			AND (created_at, id) < ($%d, $%d) 
		`, argCount, argCount+1)
		args = append(args, params.Cursor.CreatedAt, params.Cursor.LastID)
		argCount++
	}

	query += `
		ORDER BY created_at DESC, id DESC
		LIMIT $` + fmt.Sprintf("%d", argCount+1)
	args = append(args, params.PageSize)

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		logger.Error("Failed to query books", slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var books []*library_models.Book
	for rows.Next() {
		var book library_models.Book
		err = rows.Scan(&book.ID, &book.ClientID, &book.Title, &book.Author, &book.Description, &book.SubjectID, &book.ClassID, &book.CreatedAt)
		if err != nil {
			logger.Error("Failed to scan book row", slog.Any("error", err))
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		books = append(books, &book)
	}

	if err = rows.Err(); err != nil {
		logger.Error("Error during rows iteration", slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	logger.Debug("Books retrieved", slog.Int("count", len(books)))
	return books, nil
}
