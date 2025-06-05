package attachments_repository

import (
	"backend/pkg/models/library"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"log/slog"
)

func (r *Repository) DeleteAttachment(ctx context.Context, fileId string) error {
	const op = "repository.Library.Attachments.Delete"
	logger := r.logger.With(slog.String("op", op), slog.String("file_id", fileId))

	const query = `DELETE FROM attachments WHERE file_id = $1 RETURNING book_id`
	var deletedBookID int64
	err := r.db.QueryRow(ctx, query, fileId).Scan(&deletedBookID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logger.Warn("No active attachment found to delete")
			return library_models.ErrAttachmentNotFound
		}
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			logger.Error("Database error", slog.String("code", pgErr.Code), slog.String("message", pgErr.Message))
		} else {
			logger.Error("Unexpected error", slog.String("error", err.Error()))
		}
		return fmt.Errorf("failed to delete attachment: %w", err)
	}

	logger.Info("Attachment deleted", slog.Int64("book_id", deletedBookID))
	return nil
}
