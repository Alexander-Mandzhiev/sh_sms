package attachments_repository

import (
	"backend/pkg/models/library"
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) ListByBook(ctx context.Context, bookID int64) ([]*library_models.Attachment, error) {
	const op = "repository.Library.Attachments.ListByBook"
	logger := r.logger.With(slog.String("op", op), slog.Int64("book_id", bookID))
	logger.Debug("Listing attachments from database")

	query := `SELECT book_id, format, file_id, created_at, updated_at
        FROM attachments
        WHERE book_id = $1`

	rows, err := r.db.Query(ctx, query, bookID)
	if err != nil {
		logger.Error("Database query error", "error", err)
		return nil, fmt.Errorf("failed to list attachments: %w", err)
	}
	defer rows.Close()

	attachments := make([]*library_models.Attachment, 0)
	for rows.Next() {
		var a library_models.Attachment
		err = rows.Scan(&a.BookID, &a.Format, &a.FileID, &a.CreatedAt, &a.UpdatedAt)
		if err != nil {
			logger.Error("Failed to scan row", "error", err)
			return nil, fmt.Errorf("failed to scan attachment: %w", err)
		}
		attachments = append(attachments, &a)
	}

	if err = rows.Err(); err != nil {
		logger.Error("Rows iteration error", "error", err)
		return nil, fmt.Errorf("rows iteration failed: %w", err)
	}

	logger.Debug("Attachments listed from database", slog.Int("count", len(attachments)))
	return attachments, nil
}
