package attachments_repository

import (
	library_models "backend/pkg/models/library"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"log/slog"
)

func (r *Repository) GetAttachment(ctx context.Context, bookID int64, format string) (*library_models.Attachment, error) {
	const op = "repository.Library.Attachments.Get"
	logger := r.logger.With(slog.String("op", op), slog.Int64("book_id", bookID), slog.String("format", format))
	logger.Debug("Getting attachment from database")

	var attachment library_models.Attachment
	var deletedAt pgtype.Timestamp

	const query = `SELECT book_id, format, file_url, deleted_at, created_at, updated_at FROM attachments WHERE book_id = $1 AND format = $2`
	err := r.db.QueryRow(ctx, query, bookID, format).Scan(&attachment.BookID, &attachment.Format, &attachment.FileURL, &deletedAt, &attachment.CreatedAt, &attachment.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logger.Warn("Attachment not found")
			return nil, library_models.ErrAttachmentNotFound
		}
		logger.Error("Database error", "error", err)
		return nil, fmt.Errorf("failed to get attachment: %w", err)
	}

	if deletedAt.Valid {
		attachment.DeletedAt = &deletedAt.Time
	}

	logger.Debug("Attachment retrieved from database")
	return &attachment, nil
}
