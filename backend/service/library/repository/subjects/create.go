package subjects_repository

import (
	"backend/pkg/models/library"
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5/pgconn"
)

func (r *Repository) CreateSubject(ctx context.Context, subject *library_models.Subject) (int32, error) {
	const op = "repository.Library.Subjects.CreateSubject"
	logger := r.logger.With(slog.String("op", op))

	query := `
        INSERT INTO subjects (name) 
        VALUES ($1)
        RETURNING id
    `

	var id int32
	err := r.db.QueryRow(ctx, query, subject.Name).Scan(&id)
	if err != nil {
		logger.Error("failed to create subject", slog.String("error", err.Error()), slog.String("name", subject.Name))
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return 0, library_models.ErrDuplicateName
		}

		return 0, fmt.Errorf("failed to create subject: %w", err)
	}

	logger.Debug("subject created in database", slog.Int("id", int(id)), slog.String("name", subject.Name))
	return id, nil
}
