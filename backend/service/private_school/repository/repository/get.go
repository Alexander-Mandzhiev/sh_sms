package students_repository

import (
	"backend/pkg/models/students"
	"backend/pkg/utils"
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (r *Repository) GetStudent(ctx context.Context, id, clientID uuid.UUID) (*students_models.Student, error) {
	const op = "repository.PrivateSchool.StudentsRepository.GetStudent"
	logger := r.logger.With(slog.String("op", op), slog.String("student_id", id.String()), slog.String("client_id", clientID.String()), slog.String("trace_id", utils.TraceIDFromContext(ctx)))
	logger.Debug("fetching student")

	if ctx.Err() != nil {
		logger.Warn("context canceled before query execution")
		return nil, ctx.Err()
	}

	query := `SELECT id, client_id, full_name, contract_number, phone, email, additional_info, deleted_at, created_at, updated_at FROM students WHERE id = $1 AND client_id = $2`
	var st students_models.Student
	err := r.db.QueryRow(ctx, query, id, clientID).Scan(&st.ID, &st.ClientID, &st.FullName, &st.ContractNumber, &st.Phone, &st.Email, &st.AdditionalInfo, &st.DeletedAt, &st.CreatedAt, &st.UpdatedAt)
	if err != nil {
		return nil, r.handleGetError(err, op, logger)
	}
	if st.DeletedAt != nil {
		logger.Warn("student is soft-deleted", "deleted_at", st.DeletedAt)
		return nil, students_models.ErrStudentAlreadyDeleted
	}

	logger.Debug("student retrieved successfully")
	return &st, nil
}

func (r *Repository) handleGetError(err error, op string, logger *slog.Logger) error {
	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		logger.Warn("context error", "error", err)
		return err
	}

	if errors.Is(err, pgx.ErrNoRows) {
		logger.Warn("student not found")
		return students_models.ErrStudentNotFound
	}

	logger.Error("database error", "error", err)
	return fmt.Errorf("%w: %v", students_models.ErrGetFailed, err)
}
