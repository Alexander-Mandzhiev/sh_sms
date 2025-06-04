package attachments_repository

import (
	"backend/pkg/models/library"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
	"log/slog"
)

func (r *Repository) RestoreAttachment(ctx context.Context, bookID int64, format string) (*library_models.Attachment, error) {
	const op = "repository.Library.Attachments.Restore"
	logger := r.logger.With(slog.String("op", op), slog.Int64("book_id", bookID), slog.String("format", format))

	const query = `UPDATE attachments SET deleted_at = NULL, updated_at = NOW() WHERE book_id = $1 AND format = $2 AND deleted_at IS NOT NULL
                   RETURNING book_id, format, file_url, created_at, updated_at, deleted_at`

	var attachment library_models.Attachment
	var deletedAt pgtype.Timestamp

	err := r.db.QueryRow(ctx, query, bookID, format).Scan(
		&attachment.BookID,
		&attachment.Format,
		&attachment.FileURL,
		&attachment.CreatedAt,
		&attachment.UpdatedAt,
		&deletedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logger.Warn("No deleted attachment found to restore")
			return nil, library_models.ErrAttachmentNotFound
		}

		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return nil, library_models.ErrAttachmentRestoreConflict
			}
			logger.Error("Database error", slog.String("code", pgErr.Code), slog.String("message", pgErr.Message))
		} else {
			logger.Error("Unexpected error", slog.String("error", err.Error()))
		}

		return nil, fmt.Errorf("failed to restore attachment: %w", err)
	}

	attachment.DeletedAt = nil
	logger.Info("Attachment restored successfully")
	return &attachment, nil
}
