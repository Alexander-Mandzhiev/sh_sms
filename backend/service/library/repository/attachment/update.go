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

func (r *Repository) UpdateAttachment(ctx context.Context, attachment *library_models.Attachment) error {
	const op = "repository.Library.Attachments.Update"
	logger := r.logger.With(slog.String("op", op), slog.Int64("book_id", attachment.BookID), slog.String("format", attachment.Format))
	const query = `UPDATE attachments SET file_url = $1, updated_at = NOW() WHERE book_id = $2 AND format = $3 AND deleted_at IS NULL RETURNING updated_at`
	err := r.db.QueryRow(ctx, query, attachment.FileURL, attachment.BookID, attachment.Format).Scan(&attachment.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logger.Warn("No active attachment found to update")
			return library_models.ErrAttachmentNotFound
		}
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			logger.Error("Database error", slog.String("code", pgErr.Code), slog.String("message", pgErr.Message))
		} else {
			logger.Error("Unexpected error", slog.String("error", err.Error()))
		}
		return fmt.Errorf("failed to update attachment: %w", err)
	}

	logger.Info("Attachment updated successfully")
	return nil
}
