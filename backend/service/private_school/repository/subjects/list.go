package subjects_repository

import (
	private_school_models "backend/pkg/models/private_school"
	"context"
	"fmt"
	"log/slog"
)

func (r *Repository) ListSubjects(ctx context.Context) ([]*private_school_models.Subject, error) {
	const op = "repository.PrivateSchool.Subjects.ListSubjects"
	logger := r.logger.With(slog.String("op", op))

	query := `SELECT id, name FROM subjects ORDER BY id ASC`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		logger.Error("failed to query subjects list", slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to query subjects list: %w", err)
	}
	defer rows.Close()

	var subjects []*private_school_models.Subject
	for rows.Next() {
		var subject private_school_models.Subject
		if err = rows.Scan(&subject.ID, &subject.Name); err != nil {
			logger.Error("failed to scan subject row", slog.String("error", err.Error()))
			return nil, fmt.Errorf("failed to scan subject row: %w", err)
		}
		subjects = append(subjects, &subject)
	}

	if err = rows.Err(); err != nil {
		logger.Error("error during rows iteration", slog.String("error", err.Error()))
		return nil, fmt.Errorf("error during rows iteration: %w", err)
	}

	logger.Debug("subjects listed from database", slog.Int("count", len(subjects)))
	return subjects, nil
}
