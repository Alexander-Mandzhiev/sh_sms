package subjects_repository

import (
	private_school_models "backend/pkg/models/private_school"
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) DeleteSubject(ctx context.Context, id int32) error {
	const op = "repository.PrivateSchool.Subjects.DeleteSubject"
	logger := r.logger.With(slog.String("op", op))

	query := `DELETE FROM subjects WHERE id = $1`
	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		logger.Error("failed to delete subject", slog.String("error", err.Error()), slog.Int("id", int(id)))
		return fmt.Errorf("failed to delete subject: %w", err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		logger.Debug("subject not found for deletion", slog.Int("id", int(id)))
		return private_school_models.ErrNotFoundSubjectName
	}

	logger.Debug("subject deleted from database", slog.Int("id", int(id)), slog.Int64("rows_affected", rowsAffected))
	return nil
}
