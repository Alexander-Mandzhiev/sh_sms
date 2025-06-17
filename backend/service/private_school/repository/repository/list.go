package students_repository

import (
	"backend/pkg/models/students"
	"backend/pkg/utils"
	"context"
	"fmt"
	"log/slog"

	"github.com/google/uuid"
)

func (r *Repository) ListStudents(ctx context.Context, params *students_models.ListStudentsRequest) ([]*students_models.Student, string, error) {
	const op = "repository.PrivateSchool.StudentsRepository.ListStudents"
	logger := r.logger.With(slog.String("op", op), slog.String("client_id", params.ClientID.String()), slog.String("trace_id", utils.TraceIDFromContext(ctx)))
	logger.Debug("listing students", "count", params.Count, "cursor", params.Cursor)

	if ctx.Err() != nil {
		logger.Warn("context canceled before query execution")
		return nil, "", ctx.Err()
	}

	query := `SELECT id, client_id, full_name, contract_number, phone, email, additional_info, deleted_at, created_at, updated_at
		FROM students WHERE client_id = $1 AND deleted_at IS NULL`
	args := []interface{}{params.ClientID}
	argCounter := 2

	if params.Cursor != "" {
		cursor, err := uuid.Parse(params.Cursor)
		if err != nil {
			logger.Warn("invalid cursor format", "cursor", params.Cursor, "error", err)
			return nil, "", students_models.ErrInvalidCursor
		}
		query += fmt.Sprintf(" AND id > $%d", argCounter)
		args = append(args, cursor)
		argCounter++
	}

	if params.Filter != nil {
		if params.Filter.FullName != nil {
			query += fmt.Sprintf(" AND full_name ILIKE $%d", argCounter)
			args = append(args, "%"+*params.Filter.FullName+"%")
			argCounter++
		}
		if params.Filter.ContractNumber != nil {
			query += fmt.Sprintf(" AND contract_number = $%d", argCounter)
			args = append(args, *params.Filter.ContractNumber)
			argCounter++
		}
		if params.Filter.Phone != nil {
			query += fmt.Sprintf(" AND phone = $%d", argCounter)
			args = append(args, *params.Filter.Phone)
			argCounter++
		}
		if params.Filter.Email != nil {
			query += fmt.Sprintf(" AND email = $%d", argCounter)
			args = append(args, *params.Filter.Email)
			argCounter++
		}
	}

	query += fmt.Sprintf(" ORDER BY id ASC LIMIT $%d", argCounter)
	args = append(args, params.Count)

	logger.Debug("executing query", "query", query, "args", args)

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		logger.Error("query execution failed", "error", err)
		return nil, "", fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	students := make([]*students_models.Student, 0, params.Count)

	for rows.Next() {
		var st students_models.Student
		err = rows.Scan(&st.ID, &st.ClientID, &st.FullName, &st.ContractNumber, &st.Phone, &st.Email, &st.AdditionalInfo, &st.DeletedAt, &st.CreatedAt, &st.UpdatedAt)
		if err != nil {
			logger.Error("failed to scan student row", "error", err)
			return nil, "", fmt.Errorf("failed to scan student: %w", err)
		}
		students = append(students, &st)
	}

	if err = rows.Err(); err != nil {
		logger.Error("rows iteration error", "error", err)
		return nil, "", fmt.Errorf("rows iteration error: %w", err)
	}

	var nextCursor string
	if len(students) > 0 {
		nextCursor = students[len(students)-1].ID.String()
	}

	logger.Debug("students listed", "count", len(students), "next_cursor", nextCursor)
	return students, nextCursor, nil
}
