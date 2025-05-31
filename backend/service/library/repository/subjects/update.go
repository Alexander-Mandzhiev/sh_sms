package subjects_repository

import (
	"backend/pkg/models/library"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	"log/slog"
)

func (r *Repository) UpdateSubject(ctx context.Context, subject *library_models.Subject) error {
	const op = "repository.Library.Subjects.UpdateSubject"
	logger := r.logger.With(slog.String("op", op))

	query := `
        UPDATE subjects 
        SET name = $1
        WHERE id = $2
    `

	result, err := r.db.Exec(ctx, query, subject.Name, subject.ID)
	if err != nil {
		logger.Error("failed to update subject", slog.String("error", err.Error()), slog.Int("id", int(subject.ID)), slog.String("name", subject.Name))
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return library_models.ErrDuplicateName
		}
		return fmt.Errorf("failed to update subject: %w", err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		logger.Debug("subject not found for update", slog.Int("id", int(subject.ID)))
		return library_models.ErrNotFound
	}

	logger.Debug("subject updated in database", slog.Int("id", int(subject.ID)), slog.String("name", subject.Name), slog.Int64("rows_affected", rowsAffected))
	return nil
}
