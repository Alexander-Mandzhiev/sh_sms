package subjects_repository

import (
	"backend/pkg/models/subject"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log/slog"
)

func (r *Repository) GetSubjectByID(ctx context.Context, id int32) (*subjects_models.Subject, error) {
	const op = "repository.PrivateSchool.Subjects.GetSubjectByID"
	logger := r.logger.With(slog.String("op", op))
	query := `SELECT id, name FROM subjects WHERE id = $1`
	var subject subjects_models.Subject

	err := r.db.QueryRow(ctx, query, id).Scan(&subject.ID, &subject.Name)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logger.Debug("subject not found", slog.Int("id", int(id)))
			return nil, subjects_models.ErrNotFoundSubjectName
		}
		logger.Error("failed to get subject by ID", slog.String("error", err.Error()), slog.Int("id", int(id)))
		return nil, fmt.Errorf("failed to get subject by ID: %w", err)
	}

	logger.Debug("subject found in database", slog.Int("id", int(id)), slog.String("name", subject.Name))
	return &subject, nil
}
