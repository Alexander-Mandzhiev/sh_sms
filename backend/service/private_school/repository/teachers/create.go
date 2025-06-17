package teachers_repository

import (
	"backend/pkg/models/teacher"
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *Repository) CreateTeacher(ctx context.Context, teacher *teachers_models.CreateTeacher) (*teachers_models.Teacher, error) {
	const op = "repository.PrivateSchool.Teachers.CreateTeacher"
	logger := r.logger.With(slog.String("op", op))
	logger.Debug("Creating teacher in repository")

	query := `INSERT INTO teachers (id, client_id, full_name, phone, email, additional_info) VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id, client_id, full_name, phone, email, additional_info, deleted_at, created_at, updated_at`
	var t teachers_models.Teacher
	err := r.db.QueryRow(ctx, query, teacher.ID, teacher.ClientID, teacher.FullName, teacher.Phone, teacher.Email, teacher.AdditionalInfo).
		Scan(&t.ID, &t.ClientID, &t.FullName, &t.Phone, &t.Email, &t.AdditionalInfo, &t.DeletedAt, &t.CreatedAt, &t.UpdatedAt)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				logger.Warn("duplicate data detected", "detail", pgErr.Detail, "constraint", pgErr.ConstraintName)
				return nil, teachers_models.ErrDuplicateTeacher
			}
			if pgErr.Code == "23503" {
				logger.Warn("invalid client reference", "detail", pgErr.Detail, "constraint", pgErr.ConstraintName)
				return nil, teachers_models.ErrInvalidClient
			}
		}
		if errors.Is(err, pgx.ErrNoRows) {
			logger.Warn("no teacher created but no error returned")
			return nil, teachers_models.ErrCreateFailed
		}

		logger.Error("failed to create teacher", "error", err)
		return nil, fmt.Errorf("failed to create teacher: %w", err)
	}

	logger.Info("teacher created successfully", "teacher_id", teacher.ID, "client_id", teacher.ClientID)
	return &t, nil
}
