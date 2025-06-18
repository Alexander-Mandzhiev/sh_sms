package students_repository

import (
	"backend/pkg/models/students"
	"backend/pkg/utils"
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *Repository) UpdateStudent(ctx context.Context, updateData *students_models.UpdateStudent) (*students_models.Student, error) {
	const op = "repository.PrivateSchool.StudentsRepository.UpdateStudent"
	logger := r.logger.With(slog.String("op", op), slog.String("student_id", updateData.ID.String()), slog.String("client_id", updateData.ClientID.String()), slog.String("trace_id", utils.TraceIDFromContext(ctx)))
	logger.Debug("updating student")

	if ctx.Err() != nil {
		logger.Warn("request canceled before execution")
		return nil, ctx.Err()
	}

	query := `UPDATE students SET
            full_name = COALESCE(NULLIF($1, ''), full_name),
            contract_number = COALESCE(NULLIF($2, ''), contract_number),
            phone = COALESCE(NULLIF($3, ''), phone),
            email = COALESCE(NULLIF($4, ''), email),
            additional_info = COALESCE(NULLIF($5, ''), additional_info),
            updated_at = CURRENT_TIMESTAMP
        WHERE id = $6 AND client_id = $7 AND deleted_at IS NULL
		RETURNING id, client_id, full_name, contract_number, phone, email, additional_info, deleted_at, created_at, updated_at`

	args := []interface{}{
		utils.NilIfEmpty(updateData.FullName),
		utils.NilIfEmpty(updateData.ContractNumber),
		utils.NilIfEmpty(updateData.Phone),
		utils.NilIfEmpty(updateData.Email),
		utils.NilIfEmpty(updateData.AdditionalInfo),
		updateData.ID,
		updateData.ClientID,
	}

	var student students_models.Student
	err := r.db.QueryRow(ctx, query, args...).Scan(&student.ID, &student.ClientID, &student.FullName, &student.ContractNumber,
		&student.Phone, &student.Email, &student.AdditionalInfo, &student.DeletedAt, &student.CreatedAt, &student.UpdatedAt)

	if err != nil {
		return handleUpdateError(err, logger)
	}

	logger.Info("student updated successfully", "student_id", student.ID)
	return &student, nil
}

func handleUpdateError(err error, logger *slog.Logger) (*students_models.Student, error) {
	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		return nil, err
	}

	if errors.Is(err, pgx.ErrNoRows) {
		logger.Warn("student not found or already deleted")
		return nil, students_models.ErrStudentNotFound
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		if pgErr.Code == "23505" && pgErr.ConstraintName == "uniq_student_contract_active" {
			logger.Warn("duplicate contract detected")
			return nil, students_models.ErrDuplicateContract
		}

		if pgErr.Code == "23503" && pgErr.ConstraintName == "fk_client" {
			logger.Warn("invalid client reference")
			return nil, students_models.ErrInvalidClientID
		}
	}

	logger.Error("database update error", "error", err)
	return nil, fmt.Errorf("%w: %v", students_models.ErrUpdateFailed, err)
}
