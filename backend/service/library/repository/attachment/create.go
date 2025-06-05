package attachments_repository

import (
	"backend/pkg/models/library"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	"log/slog"
)

func (r *Repository) CreateAttachment(ctx context.Context, attachment *library_models.Attachment) error {
	const op = "repository.Library.Attachments.Create"
	logger := r.logger.With(slog.String("op", op), slog.Int64("book_id", attachment.BookID), slog.String("format", attachment.Format))
	const query = `INSERT INTO attachments (book_id, format, file_id) VALUES ($1, $2, $3) RETURNING created_at, updated_at`
	err := r.db.QueryRow(ctx, query, attachment.BookID, attachment.Format, attachment.FileID).Scan(&attachment.CreatedAt, &attachment.UpdatedAt)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				logger.Warn("Attachment already exists", slog.String("constraint", pgErr.ConstraintName))
				return library_models.ErrAttachmentAlreadyExists
			}
			logger.Error("Database error", slog.String("code", pgErr.Code), slog.String("message", pgErr.Message))
		}
		return fmt.Errorf("failed to create attachment: %w", err)
	}

	logger.Info("Attachment created successfully")
	return nil
}
