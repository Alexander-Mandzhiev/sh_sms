package students_repository

import (
	"backend/pkg/models/students"
	"backend/pkg/utils"
	"context"
	"fmt"
	"log/slog"
	"strings"
)

func (r *Repository) ListStudents(ctx context.Context, params *students_models.ListStudentsRequest) ([]*students_models.Student, *students_models.Cursor, error) {
	const op = "repository.StudentsRepository.ListStudents"
	logger := r.logger.With(slog.String("op", op), slog.String("client_id", params.ClientID.String()), slog.String("trace_id", utils.TraceIDFromContext(ctx)))
	logger.Debug("listing students", "count", params.Count, "filter", params.Filter)

	if ctx.Err() != nil {
		logger.Warn("context canceled before execution")
		return nil, nil, ctx.Err()
	}

	query := `SELECT id, client_id, full_name, contract_number, phone, email, additional_info, deleted_at, created_at, updated_at 
        FROM students WHERE client_id = $1 AND deleted_at IS NULL`
	args := []interface{}{params.ClientID}
	argPos := 2

	if params.Cursor != nil {
		query += fmt.Sprintf(` 
            AND (created_at, id) > ($%d, $%d) 
        `, argPos, argPos+1)
		args = append(args, params.Cursor.CreatedAt, params.Cursor.ID)
		argPos += 2
	}

	if params.Filter != "" {
		if len(params.Filter) > 100 {
			logger.Warn("filter too long", "length", len(params.Filter))
			return nil, nil, students_models.ErrFilterTooLong
		}

		filter := strings.ReplaceAll(params.Filter, "%", "\\%")
		filter = strings.ReplaceAll(filter, "_", "\\_")
		search := "%" + strings.ToLower(filter) + "%"

		query += fmt.Sprintf(` 
            AND (
                LOWER(full_name) LIKE $%d OR
                LOWER(contract_number) LIKE $%d OR
                LOWER(phone) LIKE $%d OR
                LOWER(email) LIKE $%d
            )
        `, argPos, argPos+1, argPos+2, argPos+3)
		args = append(args, search, search, search, search)
		argPos += 4
	}

	query += fmt.Sprintf(` 
        ORDER BY created_at ASC, id ASC 
        LIMIT $%d 
    `, argPos)
	args = append(args, params.Count+1)

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		logger.Error("query execution failed", "error", err, "query", query)
		return nil, nil, fmt.Errorf("%w: %v", students_models.ErrListFailed, err)
	}
	defer rows.Close()

	var students []*students_models.Student
	for rows.Next() {
		var s students_models.Student
		if err = rows.Scan(&s.ID, &s.ClientID, &s.FullName, &s.ContractNumber, &s.Phone, &s.Email, &s.AdditionalInfo, &s.DeletedAt, &s.CreatedAt, &s.UpdatedAt); err != nil {
			logger.Error("failed to scan student", "error", err)
			return nil, nil, fmt.Errorf("%w: scan failed: %v", students_models.ErrListFailed, err)
		}
		students = append(students, &s)
	}

	if err = rows.Err(); err != nil {
		logger.Error("rows iteration error", "error", err)
		return nil, nil, fmt.Errorf("%w: iteration failed: %v", students_models.ErrListFailed, err)
	}

	var nextCursor *students_models.Cursor
	if len(students) == int(params.Count)+1 {
		lastStudent := students[len(students)-1]
		students = students[:len(students)-1]

		nextCursor = &students_models.Cursor{
			ID:        lastStudent.ID,
			CreatedAt: lastStudent.CreatedAt,
		}
		logger.Debug("next cursor created", "cursor_id", nextCursor.ID, "cursor_at", nextCursor.CreatedAt)
	}

	logger.Debug("students listed", "count", len(students))
	return students, nextCursor, nil
}
