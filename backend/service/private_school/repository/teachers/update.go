package teachers_repository

import (
	"backend/pkg/models/private_school"
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *Repository) UpdateTeacher(ctx context.Context, update *private_school_models.UpdateTeacher) (*private_school_models.Teacher, error) {
	const op = "repository.PrivateSchool.Teachers.UpdateTeacher"
	logger := r.logger.With(slog.String("op", op))
	logger.Debug("Updating teacher in repository")

	query := `UPDATE teachers SET
		full_name = COALESCE(NULLIF($3, ''), full_name),
		phone = COALESCE(NULLIF($4, ''), phone),
		email = COALESCE(NULLIF($5, ''), email),
		additional_info = COALESCE(NULLIF($6, ''), additional_info),
		updated_at = NOW()
	WHERE id = $1 AND client_id = $2
	RETURNING id, client_id, full_name, phone, email, additional_info, deleted_at, created_at, updated_at`

	var t private_school_models.Teacher
	err := r.db.QueryRow(ctx, query, update.ID, update.ClientID, update.FullName, update.Phone, update.Email, update.AdditionalInfo).
		Scan(&t.ID, &t.ClientID, &t.FullName, &t.Phone, &t.Email, &t.AdditionalInfo, &t.DeletedAt, &t.CreatedAt, &t.UpdatedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logger.Warn("teacher not found", "teacher_id", update.ID, "client_id", update.ClientID)
			return nil, private_school_models.ErrTeacherNotFound
		}
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			logger.Warn("duplicate data detected", "detail", pgErr.Detail, "constraint", pgErr.ConstraintName)
			return nil, private_school_models.ErrDuplicateTeacher
		}

		logger.Error("failed to update teacher", "error", err)
		return nil, fmt.Errorf("failed to update teacher: %w", err)
	}

	logger.Info("teacher updated successfully", "teacher_id", update.ID, "client_id", update.ClientID)
	return &t, nil
}
