package students_repository

import (
	"backend/pkg/models/students"
	"backend/pkg/utils"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"log/slog"
)

func (r *Repository) CreateStudent(ctx context.Context, s *students_models.CreateStudent) (*students_models.Student, error) {
	const op = "repository.PrivateSchool.StudentsRepository.CreateStudent"
	logger := r.logger.With(slog.String("op", op), slog.String("client_id", s.ClientID.String()), slog.String("contract", s.ContractNumber), slog.String("trace_id", utils.TraceIDFromContext(ctx)))
	logger.Debug("creating student in repository")

	newID := uuid.New()
	query := `INSERT INTO students (id, client_id, full_name, contract_number, phone, email, additional_info) VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, client_id, full_name, contract_number, phone, email, additional_info, deleted_at, created_at, updated_at`

	params := []interface{}{newID, s.ClientID, s.FullName, s.ContractNumber, s.Phone, s.Email, s.AdditionalInfo}

	if ctx.Err() != nil {
		logger.Warn("context canceled before query execution")
		return nil, ctx.Err()
	}

	var student students_models.Student
	err := r.db.QueryRow(ctx, query, params...).Scan(&student.ID, &student.ClientID, &student.FullName, &student.ContractNumber,
		&student.Phone, &student.Email, &student.AdditionalInfo, &student.DeletedAt, &student.CreatedAt, &student.UpdatedAt)
	if err != nil {
		return nil, r.handleCreateError(err, op, query, params, logger)
	}

	logger.Info("student created", "student_id", student.ID)
	return &student, nil
}

func (r *Repository) handleCreateError(err error, op, query string, params []interface{}, logger *slog.Logger) error {
	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		logger.Warn("context error", "error", err)
		return err
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch {
		case pgErr.Code == "23505" && pgErr.ConstraintName == "uniq_student_contract_active":
			logger.Warn("duplicate contract detected", "detail", pgErr.Detail)
			return fmt.Errorf("%w: %s", students_models.ErrDuplicateContract, pgErr.Detail)

		case pgErr.Code == "23503" && pgErr.ConstraintName == "fk_client":
			logger.Warn("invalid client reference", "detail", pgErr.Detail)
			return students_models.ErrInvalidClientID

		default:
			logger.Warn("database constraint violation", "code", pgErr.Code, "constraint", pgErr.ConstraintName, "detail", pgErr.Detail)
			return fmt.Errorf("database constraint violation: %w", err)
		}
	}

	if errors.Is(err, pgx.ErrNoRows) {
		logger.Warn("no student created after insert")
		return students_models.ErrCreateFailed
	}
	logger.Error("database operation failed", "error", err, "query", query, "params", params)
	return fmt.Errorf("%s: %w", op, err)
}
